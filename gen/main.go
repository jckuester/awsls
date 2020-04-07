// +build codegen

package main

import (
	"github.com/apex/log"
	"github.com/jckuester/terratools/gen/terraform"
)

func main() {
	resourceTypes := terraform.ResourceTypes()

	err := terraform.WriteResourceTypes("../resource", resourceTypes)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Infof("Generated list of Terraform AWS resource types: %d", len(resourceTypes))
}
