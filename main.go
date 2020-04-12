package main

import (
	"os"

	"github.com/jckuester/awsls/resource"

	"github.com/jckuester/awsls/aws"
)

func main() {
	client := aws.NewClient()

	if os.Args[1] == "*" {
		resource.ListResources(client)
		return
	}

	resource.ListResourcesByType(client, os.Args[1])
}
