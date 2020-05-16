package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	stdlog "log"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/fatih/color"
	"github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/internal"
	"github.com/jckuester/awsls/resource"
	"github.com/jckuester/terradozer/pkg/provider"
	terradozerRes "github.com/jckuester/terradozer/pkg/resource"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func main() {
	os.Exit(mainExitCode())
}

func mainExitCode() int {
	var logDebug bool
	var profile string
	var region string
	var attribute string
	var version bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flags.Usage = func() {
		printHelp(flags)
	}

	flags.BoolVar(&logDebug, "debug", false, "Enable debug logging")
	flags.StringVar(&profile, "profile", "", "Use a specific named profile from your AWS credential file")
	flags.StringVar(&region, "region", "", "The region to list resources in")
	flags.StringVar(&attribute, "attribute", "", "An attribute to show for each resources")
	flags.BoolVar(&version, "version", false, "Show application version")

	_ = flags.Parse(os.Args[1:])
	args := flags.Args()

	// discard TRACE logs of GRPCProvider
	stdlog.SetOutput(ioutil.Discard)

	start := time.Now()

	fmt.Println()
	defer fmt.Println()
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

	if profile != "" {
		err := os.Setenv("AWS_PROFILE", profile)
		if err != nil {
			log.WithError(err).Error("failed to set AWS profile")
			return 1
		}
	}
	if region != "" {
		err := os.Setenv("AWS_DEFAULT_REGION", region)
		if err != nil {
			log.WithError(err).Error("failed to set AWS region")
			return 1
		}
	}

	provider, err := provider.Init("aws", 10*time.Second)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("\nError:️ failed to initialize Terraform AWS provider: %s\n", err))
		return 1
	}

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

	client, err := aws.NewClient()
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))

		return 1
	}

	for _, rType := range matchedTypes {
		resources, err := aws.ListResourcesByType(client, rType)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))
		}

		resourcesWithStates := GetStates(resources, provider)

		printResources(resourcesWithStates, attribute)
	}

	log.Debugf("List completed: %s", time.Since(start))

	return 0
}

// GetStates fetches the Terraform state for each resource via the Terraform AWS Provider.
func GetStates(resources []aws.Resource, provider *provider.TerraformProvider) []aws.Resource {
	var wg sync.WaitGroup

	start := time.Now()

	sem := internal.NewSemaphore(10)
	for i := range resources {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Acquire a semaphore so that we can limit concurrency
			sem.Acquire()
			defer sem.Release()

			r := &resources[i]
			r.Resource = terradozerRes.New(r.Type, r.ID, provider)

			log.Debugf("Update state: %s", r.ID)

			err := r.UpdateState()
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))
			}
		}(i)
	}

	// Wait for all updates to complete
	wg.Wait()

	log.Debugf("Completed fetching states: %s", time.Since(start))

	return resources
}

func printResources(resources []aws.Resource, attribute string) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	for _, r := range resources {
		fmt.Fprintf(w, "%s\t%s\t%s\t", r.Type, r.ID, r.Region)

		if r.CreatedAt != nil {
			fmt.Fprintf(w, "%s\t", r.CreatedAt.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Fprint(w, "-\t")
		}

		v, err := GetAttribute(attribute, r.State())
		if err != nil {
			log.WithError(err).Debug("failed to get attribute")
			v = "-"
		}

		fmt.Fprintf(w, "%s\t\n", v)
	}

	w.Flush()
}

func GetAttribute(attrName string, state *cty.Value) (string, error) {
	if state == nil {
		return "", fmt.Errorf("state is nil pointer")
	}

	if state.IsNull() {
		return "", fmt.Errorf("state is nil value")
	}

	if !state.IsKnown() {
		return "", fmt.Errorf("state is not known")
	}

	if !state.IsWhollyKnown() {
		return "", fmt.Errorf("state is not wholly known")
	}

	if !state.CanIterateElements() {
		return "", fmt.Errorf("cannot iterate: %s", *state)
	}

	attrValue, ok := state.AsValueMap()[attrName]
	if !ok {
		return "", fmt.Errorf("attribute not found: %s", attrName)
	}

	switch attrValue.Type() {
	case cty.Bool:
		var v bool
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		return strconv.FormatBool(v), nil

	case cty.Number:
		var v int64
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		return strconv.FormatInt(v, 10), nil

	case cty.String:
		var v string
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		return v, nil

	case cty.Map(cty.String):
		var v map[string]string
		err := gocty.FromCtyValue(attrValue, &v)
		if err != nil {
			return "", err
		}

		if len(v) > 0 {
			var tagList []string

			for k, v := range v {
				tagList = append(tagList, fmt.Sprintf("%s=%s", k, v))
			}

			return strings.Join(tagList, ","), nil
		}

		return "-", nil

	default:
		return "", fmt.Errorf("currently unhandled type: %s", attrValue.Type())
	}
}

func printHelp(fs *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "\n"+strings.TrimSpace(help)+"\n")
	fs.PrintDefaults()
	fmt.Println()
}

const help = `
awsls - list AWS resources.

USAGE:
  $ awsls [flags] [<resource_type glob pattern>]

FLAGS:
`
