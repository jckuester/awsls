package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	stdlog "log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/jckuester/terradozer/pkg/provider"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	aws_ssmhelpers "github.com/disneystreaming/go-ssmhelpers/aws"
	"github.com/fatih/color"
	awsls "github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/internal"
	"github.com/jckuester/awsls/resource"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
	flag "github.com/spf13/pflag"
)

func main() {
	os.Exit(mainExitCode())
}

type UpdatedResources struct {
	Resources []awsls.Resource
	Errors    []error
}

func mainExitCode() int {
	var logDebug bool
	var allProfilesFlag bool
	var profiles internal.CommaSeparatedListFlag
	var regions internal.CommaSeparatedListFlag
	var attributes internal.CommaSeparatedListFlag
	var version bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flags.Usage = func() {
		printHelp(flags)
	}

	flags.BoolVar(&logDebug, "debug", false, "Enable debug logging")
	flags.VarP(&profiles, "profiles", "p", "Comma-separated list of named AWS profiles for accounts to list resources in")
	flags.BoolVar(&allProfilesFlag, "all-profiles", false, "List resources for all profiles in ~/.aws/config")
	flags.VarP(&regions, "regions", "r", "Comma-separated list of regions to list resources in")
	flags.VarP(&attributes, "attributes", "a", "Comma-separated list of attributes to show for each resource")
	flags.BoolVar(&version, "version", false, "Show application version")

	_ = flags.Parse(os.Args[1:])
	args := flags.Args()

	fmt.Println()
	defer fmt.Println()

	// discard TRACE logs of GRPCProvider
	stdlog.SetOutput(ioutil.Discard)

	log.SetHandler(cli.Default)

	if logDebug {
		log.SetLevel(log.DebugLevel)
	}

	if version {
		fmt.Println(internal.BuildVersionString())
		return 0
	}

	if len(args) == 0 {
		fmt.Fprint(os.Stderr, color.RedString("Error: resource type glob pattern expected\n"))
		printHelp(flags)

		return 1
	}

	if profiles != nil && allProfilesFlag == true {
		fmt.Fprint(os.Stderr, color.RedString("Error:️ --profiles and --all-profiles flag cannot be used together\n"))
		printHelp(flags)

		return 1
	}

	resourceTypePattern := args[0]
	matchedTypes, err := resource.MatchSupportedTypes(resourceTypePattern)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("Error: invalid glob pattern: %s\n", resourceTypePattern))

		return 1
	}

	if len(matchedTypes) == 0 {
		fmt.Fprint(os.Stderr, color.RedString("Error: no resource type found: %s\n", resourceTypePattern))
		return 1
	}

	if profiles == nil && allProfilesFlag == false {
		env, ok := os.LookupEnv("AWS_PROFILE")
		if ok {
			profiles = []string{env}
		}
	}

	if allProfilesFlag {
		var awsConfigPath []string
		awsConfigFileEnv, ok := os.LookupEnv("AWS_CONFIG_FILE")
		if ok {
			awsConfigPath = []string{awsConfigFileEnv}
		}

		profilesFromConfig, err := aws_ssmhelpers.GetAWSProfiles(awsConfigPath...)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString("Error: failed to load all profiles: %s\n", err))
			return 1
		}

		if profilesFromConfig == nil {
			fmt.Fprint(os.Stderr, color.RedString("Error: no profiles found in ~/.aws/config\n"))
			return 1
		}

		profiles = profilesFromConfig
	}

	clients, err := aws.NewClientPool(profiles, regions)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("\nError: %s\n", err))

		return 1
	}

	// suppress provider debug and info logs
	log.SetLevel(log.ErrorLevel)

	clientKeys := make([]aws.ClientKey, 0, len(clients))
	for k := range clients {
		clientKeys = append(clientKeys, k)
	}

	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, ignoreSignals...)
	signal.Notify(signalCh, forwardSignals...)
	defer func() {
		signal.Stop(signalCh)
		cancel()
	}()
	go func() {
		select {
		case <-signalCh:
			cancel()
		case <-ctx.Done():
		}
	}()

	if logDebug {
		log.SetLevel(log.DebugLevel)
	}

	// initialize a Terraform AWS provider for each AWS client with a matching config
	providers, err := terraform.NewProviderPool(ctx, clientKeys, "v3.16.0", "~/.awsls", 10*time.Second)
	if err != nil {
		if !errors.Is(err, context.Canceled) {
			fmt.Fprint(os.Stderr, color.RedString("\nError: %s\n", err))
		}
		return 1
	}
	defer func() {
		for _, p := range providers {
			_ = p.Close()
		}
	}()

	for _, rType := range matchedTypes {
		// any provider here is sufficient to check if a resource type has attributes
		p := providers[clientKeys[0]]

		hasAttrs, err := resource.HasAttributes(attributes, rType, &p)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString("Error: failed to check if resource type has attribute: "+
				"%s\n", err))
			return 1
		}

		var resources []awsls.Resource

		resourcesCh := make(chan UpdatedResources, 1)
		go func() { resourcesCh <- listResourcesOfType(rType, hasAttrs, clients, providers) }()
		select {
		case <-ctx.Done():
			return 1
		case result := <-resourcesCh:
			resources = result.Resources

			for _, err := range result.Errors {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))
			}
		}

		if len(resources) == 0 {
			continue
		}

		printResources(resources, hasAttrs, attributes)
	}

	return 0
}

// listResourcesOfType lists resources of given resource type in parallel for multiple accounts and regions.
func listResourcesOfType(rType string, hasAttrs map[string]bool, clients map[aws.ClientKey]awsls.Client,
	providers map[aws.ClientKey]provider.TerraformProvider) UpdatedResources {
	var wg sync.WaitGroup
	sem := internal.NewSemaphore(5)

	resources := terraform.ResourcesThreadSafe{
		Resources: []awsls.Resource{},
	}

	for key, client := range clients {
		log.WithFields(log.Fields{
			"type":    rType,
			"region":  key.Region,
			"profile": key.Profile,
			"time":    time.Now().Format("04:05.000"),
		}).Debugf("start listing resources")

		wg.Add(1)

		go func(client awsls.Client) {
			defer wg.Done()

			// Acquire a semaphore so that we can limit concurrency
			sem.Acquire()
			defer sem.Release()

			err := client.SetAccountID()
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))
				return
			}

			res, err := awsls.ListResourcesByType(&client, rType)
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))
				return
			}

			if len(hasAttrs) > 0 {
				// for performance reasons:
				// only fetch state if some attributes need to be displayed for this resource type
				updatesRes, errs := terraform.UpdateStates(res, providers)
				res = updatesRes

				resources.Lock()
				resources.Errors = append(resources.Errors, errs...)
				resources.Unlock()
			}

			resources.Lock()
			resources.Resources = append(resources.Resources, res...)
			resources.Unlock()
		}(client)
	}

	// Wait until listing resources of this type completes for every account and region
	wg.Wait()

	return UpdatedResources{resources.Resources, resources.Errors}
}

func printResources(resources []awsls.Resource, hasAttrs map[string]bool, attributes []string) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	printHeader(w, attributes)

	for _, r := range resources {
		profile := `N/A`
		if r.Profile != "" {
			profile = r.Profile
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s", r.Type, r.ID, profile, r.Region)

		if r.CreatedAt != nil {
			fmt.Fprintf(w, "\t%s", r.CreatedAt.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Fprint(w, "\tN/A")
		}

		for _, attr := range attributes {
			v := "N/A"

			_, ok := hasAttrs[attr]
			if ok {
				var err error
				v, err = resource.GetAttribute(attr, &r)
				if err != nil {
					log.WithFields(log.Fields{
						"type": r.Type,
						"id":   r.ID}).WithError(err).Debug("failed to get attribute")
					v = "error"
				}
			}

			fmt.Fprintf(w, "\t%s", v)
		}

		fmt.Fprintf(w, "\t\n")
	}

	w.Flush()
	fmt.Println()
}

func printHeader(w *tabwriter.Writer, attributes []string) {
	fmt.Fprintf(w, "TYPE\tID\tPROFILE\tREGION\tCREATED")

	for _, attribute := range attributes {
		fmt.Fprintf(w, "\t%s", strings.ToUpper(attribute))
	}

	fmt.Fprintf(w, "\t\n")
}

func printHelp(fs *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "\n"+strings.TrimSpace(help)+"\n")
	fs.PrintDefaults()
}

const help = `
awsls - list AWS resources.

USAGE:
  $ awsls [flags] <resource_type glob pattern>

FLAGS:
`
