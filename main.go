package main

import (
	"fmt"
	"io/ioutil"
	stdlog "log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/jckuester/awsls/util"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	aws_ssmhelpers "github.com/disneystreaming/go-ssmhelpers/aws"
	"github.com/fatih/color"
	"github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/internal"
	"github.com/jckuester/awsls/resource"
	flag "github.com/spf13/pflag"
)

func main() {
	os.Exit(mainExitCode())
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

		if profilesFromConfig != nil {
			profiles = profilesFromConfig
		}

	}

	clients, err := util.NewAWSClientPool(profiles, regions)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("\nError: %s\n", err))

		return 1
	}

	// suppress provider debug and info logs
	log.SetLevel(log.ErrorLevel)

	clientKeys := make([]util.AWSClientKey, 0, len(clients))
	for k := range clients {
		clientKeys = append(clientKeys, k)
	}

	// initialize a Terraform AWS provider for each AWS client with a matching config
	providers, err := util.NewProviderPool(clientKeys)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("\nError: %s\n", err))

		return 1
	}

	if logDebug {
		log.SetLevel(log.DebugLevel)
	}

	for _, rType := range matchedTypes {
		var resources []aws.Resource
		var hasAttrs map[string]bool

		for key, client := range clients {
			err := client.SetAccountID()
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))

				return 1
			}

			res, err := aws.ListResourcesByType(&client, rType)
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))

				continue
			}

			provider := providers[key]

			hasAttrs, err = resource.HasAttributes(attributes, rType, &provider)
			if err != nil {
				fmt.Fprint(os.Stderr, color.RedString("Error: failed to check if resource type has attribute: "+
					"%s\n", err))

				continue
			}

			if len(hasAttrs) > 0 {
				// for performance reasons:
				// only fetch state if some attributes need to be displayed for this resource type
				res = resource.GetStates(res, &provider)
			}

			resources = append(resources, res...)
		}

		if len(resources) == 0 {
			continue
		}

		printResources(resources, hasAttrs, attributes)
	}

	return 0
}

func printResources(resources []aws.Resource, hasAttrs map[string]bool, attributes []string) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	printHeader(w, attributes)

	for _, r := range resources {
		profile := `N\A`
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
