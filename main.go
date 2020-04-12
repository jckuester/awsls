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
	client := aws.NewClient()

	if os.Args[1] == "*" {
		resource.ListResources(client)
		return 1
	}

	resourceType := os.Args[1]

	if !resource.IsResourceType(resourceType) {
		fmt.Fprint(os.Stderr, color.RedString("Error: this is not a valid Terraform AWS resource type: %s\n",
			resourceType))
		return 1
	}

	err := resource.ListResourcesByType(client, resourceType)
	if err != nil {
		fmt.Fprint(os.Stderr, color.RedString("Error: %s\n", err))
		return 1
	}

	return 0
}
