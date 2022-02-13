//go:build codegen

//go:generate go run -tags codegen generate.go
//go:generate gofmt -s -w ../resource ../aws
//go:generate goimports -w ../resource ../aws

package main

import (
	"io/ioutil"
	stdlog "log"
	"os"
	"sort"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jckuester/awsls/gen/aws"
	"github.com/jckuester/awsls/gen/terraform"
	"github.com/jckuester/awsls/gen/util"
)

const (
	outputPathAWS                = "../aws"
	outputPathResource           = "../resource"
	terraformAwsProviderRepoPath = "/home/jan/git/github.com/terraform-provider-aws"
	awsSdkGoRepoPath             = "/home/jan/git/github.com/aws-sdk-go"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	// discard TRACE logs of GRPCProvider
	stdlog.SetOutput(ioutil.Discard)

	profile := "myaccount"
	region := "us-west-2"

	_, ok := os.LookupEnv("AWS_PROFILE")
	if !ok {
		log.Infof("AWS_PROFILE not set, therefore using the following default value: %s", profile)

		err := os.Setenv("AWS_PROFILE", profile)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	_, ok = os.LookupEnv("AWS_DEFAULT_REGION")
	if !ok {
		log.Infof("AWS_DEFAULT_REGION not set, therefore using the following default value: %s", region)

		err := os.Setenv("AWS_DEFAULT_REGION", region)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	err := os.MkdirAll(outputPathResource, 0775)
	if err != nil {
		log.Fatalf("failed to create directory: %s", err)
	}

	err = os.MkdirAll(outputPathAWS, 0775)
	if err != nil {
		log.Fatalf("failed to create directory: %s", err)
	}

	resourceTypes, err := terraform.GenerateResourceTypeList(terraformAwsProviderRepoPath, outputPathResource)
	if err != nil {
		log.WithError(err).Fatal("failed to generate list of Terraform AWS resource types")
	}

	resourceFileNames, err := terraform.ResourceFileNames(terraformAwsProviderRepoPath, resourceTypes)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> filename implementing resource")
	}

	resourceServices := terraform.GenerateResourceServiceMap(terraformAwsProviderRepoPath, outputPathResource,
		resourceTypes, resourceFileNames)

	resourceIDs, err := terraform.GenerateResourceIDMap(terraformAwsProviderRepoPath, outputPathResource, resourceFileNames)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> resource ID")
	}

	resourceTypesWithTags, err := terraform.GenerateResourceTypesWithTagsList(resourceTypes, outputPathResource)
	if err != nil {
		log.WithError(err).Fatal("failed to generate list of resource type that support tags")
	}

	apis, err := aws.APIs(awsSdkGoRepoPath)
	if err != nil {
		log.WithError(err).Fatal("failed to load AWS APIs")
	}

	services := aws.ResourceTypesForAWSServices(resourceServices)

	var serviceNamesUsedByTerraform []string
	for _, s := range services {
		serviceNamesUsedByTerraform = append(serviceNamesUsedByTerraform, s.Name)
	}

	servicePkgNames := aws.ServicePkgNames(apis)

	log.Infof("AWS generatedServices covered by Terraform: %d/%d", len(services), len(servicePkgNames))

	log.Debugf("AWS generatedServices not covered by Terraform:")
	diff := util.Difference(servicePkgNames, serviceNamesUsedByTerraform)
	for _, d := range diff {
		log.Debugf("\t%s", d)
	}

	log.Warn("AWS generatedServices used by Terraform that are named differently in AWS API v2:")
	diff = util.Difference(serviceNamesUsedByTerraform, servicePkgNames)
	for _, d := range diff {
		log.Warnf("\t: %s", d)
	}

	generatedServices := aws.GenerateListFunctions(outputPathAWS, services, resourceIDs, resourceTypesWithTags, apis)

	var rTypes []aws.ResourceType
	for _, service := range generatedServices {
		for _, rType := range service.TerraformResourceTypes {
			rTypes = append(rTypes, rType)
		}
	}

	sort.Slice(rTypes, func(i, j int) bool {
		return rTypes[i].Name < rTypes[j].Name
	})

	terraform.GenerateListResourcesByTypeFunction(outputPathAWS, rTypes)
	aws.GenerateListOfSupportedResourceTypes(outputPathResource, rTypes)
	aws.GenerateReadme("..", generatedServices, len(rTypes))

	log.Infof("Generated list functions: %d", len(rTypes))
}
