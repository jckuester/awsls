// +build codegen

package main

import (
	"fmt"
	"os"

	"github.com/jckuester/terratools/gen/util"

	"github.com/jckuester/terratools/resource"

	"github.com/aws/aws-sdk-go-v2/private/model/api"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jckuester/terratools/gen/aws"
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

	globs := []string{"/home/jan/git/github.com/aws-sdk-go-v2/models/apis/*/*/api-2.json"}
	modelPaths, err := api.ExpandModelGlobPath(globs...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to glob file pattern", err)
		os.Exit(1)
	}
	modelPaths, _ = api.TrimModelServiceVersions(modelPaths)

	loader := api.Loader{
		BaseImport:          outputPath,
		KeepUnsupportedAPIs: false,
	}

	apis, err := loader.Load(modelPaths)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load API models", err)
		os.Exit(1)
	}
	if len(apis) == 0 {
		fmt.Fprintf(os.Stderr, "expected to load models, but found none")
		os.Exit(1)
	}

	delete(apis, "github.com/aws/aws-sdk-go-v2/service/importexport")

	log.Infof("AWS services covered by Terraform: %d/%d",
		len(aws.ServicesCoveredByTerraform(resource.ResourceServices)), len(aws.Services(apis)))

	log.Debugf("AWS services not covered:")
	diff := util.Difference(aws.Services(apis), aws.ServicesCoveredByTerraform(resource.ResourceServices))
	for _, d := range diff {
		log.Debugf("\t%s", d)
	}

	diff = util.Difference(aws.ServicesCoveredByTerraform(resource.ResourceServices), aws.Services(apis))
	for _, d := range diff {
		log.Warnf("\tnot in v2 API: %s", d)
	}

	err = aws.WriteClient("../aws", aws.ServicesCoveredByTerraform(resource.ResourceServices))
	if err != nil {
		log.Fatal(err.Error())
	}
}
