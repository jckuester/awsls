package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jckuester/awsls/aws"
	"github.com/jckuester/awsls/resource"
)

func main() {
	os.Exit(mainExitCode())
}

func mainExitCode() int {
	client, err := aws.NewClient()
	if err != nil {
		return 1
	}
	fmt.Println()

	if os.Args[1] == "*" {
		aws.ListResources(client)
		return 1
	}

	resourceType := os.Args[1]

	if !resource.IsResourceType(resourceType) {
		fmt.Fprint(os.Stderr, color.RedString("Error: not a valid Terraform AWS resource type: %s\n",
			resourceType))
		return 1
	}

	resources, err := aws.ListResourcesByType(client, resourceType)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))
		return 1
	}

	for _, r := range resources {
		fmt.Printf("ID:\t%s\n", r.ID)
		if len(r.Tags) > 0 {
			fmt.Println("Tags:")
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
		if r.CreatedAt != nil {
			fmt.Printf("Created at:\t%s\n", r.CreatedAt)
		}
		fmt.Println()
	}

	return 0
}
