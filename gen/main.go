// +build codegen

package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jckuester/terratools/gen/terraform"
)

const outputPath = "../resource"

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	resourceTypes := terraform.ResourceTypes()

	err := terraform.WriteResourceTypes(outputPath, resourceTypes)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.WithField("length", len(resourceTypes)).Infof("Generated list of Terraform AWS resource types")

	resourceFileNames, err := terraform.ResourceFileNames(resourceTypes)
	if err != nil {
		log.Fatal(err.Error())
	}

	resourceService := map[string]string{}
	for _, rType := range resourceTypes {
		rFileName, ok := resourceFileNames[rType]
		if !ok {
			log.Fatal("cannot find filename for resource type")
		}

		service, err := terraform.ResourceService(rType, rFileName)
		if err != nil {
			log.WithError(err).WithField("resource", rType).Warn("could not identify AWS service for resource")
			continue
		}
		resourceService[rType] = service
	}

	err = terraform.WriteResourceServices(outputPath, resourceService)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Infof("Generated map of resource type -> AWS service: %d/%d", len(resourceService), len(resourceTypes))

	resourceIDs := map[string]string{}
	for rType, fileName := range resourceFileNames {
		resourceID, err := terraform.GetResourceID(fileName)
		if err != nil {
			continue
		}
		resourceIDs[rType] = resourceID
	}

	err = terraform.WriteResourceIDs(outputPath, resourceIDs)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Infof("Generated map of resource type -> Output field name in the AWS API of the Terraform resource ID: %d",
		len(resourceIDs))
}
