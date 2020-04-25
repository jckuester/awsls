package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/fatih/color"
	"github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/internal"
	"github.com/jckuester/awsls/resource"
)

func main() {
	os.Exit(mainExitCode())
}

func mainExitCode() int {
	var logDebug bool
	var profile string
	var region string
	var version bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flags.Usage = func() {
		printHelp(flags)
	}

	flags.BoolVar(&logDebug, "debug", false, "Enable debug logging")
	flags.StringVar(&profile, "profile", "", "Use a specific named profile from your AWS credential file")
	flags.StringVar(&region, "region", "", "The region to list resources in")
	flags.BoolVar(&version, "version", false, "Show application version")

	_ = flags.Parse(os.Args[1:])
	args := flags.Args()

	log.SetHandler(cli.Default)

	fmt.Println()
	defer fmt.Println()

	if logDebug {
		log.SetLevel(log.DebugLevel)
	}

	if version {
		fmt.Println(internal.BuildVersionString())
		return 0
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

	resourceTypePattern := args[0]

	for _, rType := range resource.MatchTypes(resourceTypePattern) {
		resources, err := aws.ListResourcesByType(client, rType)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))
		}

		printResources(resources)
	}

	return 0
}

func printResources(resources []aws.Resource) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	for _, r := range resources {
		fmt.Fprintf(w, "%s\t%s\t%s\t", r.Type, r.ID, r.Region)

		if r.CreatedAt != nil {
			fmt.Fprintf(w, "%s\t", r.CreatedAt.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Fprint(w, "-\t")
		}

		if len(r.Tags) > 0 {
			var tagList []string
			for k, v := range r.Tags {
				tagList = append(tagList, fmt.Sprintf("%s=%s", k, v))
			}
			fmt.Fprint(w, strings.Join(tagList, ",")+"\t\n")
		} else {
			fmt.Fprint(w, "-\t\n")
		}
	}

	w.Flush()
}

func printHelp(fs *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "\n"+strings.TrimSpace(help)+"\n")
	fs.PrintDefaults()
	fmt.Println()
}

const help = `
awsls - listing AWS resources via CLI.

USAGE:
  $ awsls [flags] [<resource_type glob pattern>]

FLAGS:
`
