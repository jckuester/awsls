package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	stdlog "log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/fatih/color"
	"github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/internal"
	"github.com/jckuester/awsls/resource"
	"github.com/jckuester/terradozer/pkg/provider"
)

func main() {
	os.Exit(mainExitCode())
}

func mainExitCode() int {
	var logDebug bool
	var profile string
	var region string
	var attributes internal.Attributes
	var version bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flags.Usage = func() {
		printHelp(flags)
	}

	flags.BoolVar(&logDebug, "debug", false, "Enable debug logging")
	flags.StringVar(&profile, "profile", "", "Use a specific named profile from your AWS credential file")
	flags.StringVar(&region, "region", "", "The region to list resources in")
	flags.Var(&attributes, "attributes", "Comma-separated list of attributes to show for each resource")
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

	// suppress provider debug and info logs
	log.SetLevel(log.ErrorLevel)

	awsTerraformProvider, err := provider.Init("aws", "~/.awsls", 10*time.Second)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("Error:️ failed to initialize Terraform AWS provider: %s\n", err))
		return 1
	}

	if logDebug {
		log.SetLevel(log.DebugLevel)
	}

	client, err := aws.NewClient()
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("\nError: %s\n", err))

		return 1
	}

	for _, rType := range matchedTypes {
		resources, err := aws.ListResourcesByType(client, rType)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString("Error %s: %s\n", rType, err))

			continue
		}

		if len(resources) == 0 {
			continue
		}

		hasAttrs, err := resource.HasAttributes(attributes, rType, awsTerraformProvider)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString("Error: failed to check if resource type has attribute: "+
				"%s\n", err))

			continue
		}

		if len(hasAttrs) > 0 {
			// for performance reasons, only load the state if at least one attribute needs to be displayed
			// for this resource type
			resourcesWithStates := resource.GetStates(resources, awsTerraformProvider)
			printResources(resourcesWithStates, hasAttrs, attributes)

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
		fmt.Fprintf(w, "%s\t%s\t%s", r.Type, r.ID, r.Region)

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
	fmt.Fprintf(w, "TYPE\tID\tREGION\tCREATED")

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
