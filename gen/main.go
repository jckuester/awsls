// +build codegen

package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jckuester/awsls/gen/aws"
	"github.com/jckuester/awsls/gen/terraform"
	"github.com/jckuester/awsls/gen/util"
)

const (
	outputPath                   = "../resource"
	terraformAwsProviderRepoPath = "/home/jan/git/github.com/terraform-provider-aws"
	awsSdkGoRepoPath             = "/home/jan/git/github.com/aws-sdk-go-v2"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	resourceTypes, err := terraform.GenerateResourceTypeList(terraformAwsProviderRepoPath, outputPath)
	if err != nil {
		log.WithError(err).Fatal("failed to generate list of Terraform AWS resource types")
	}

	resourceFileNames, err := terraform.ResourceFileNames(terraformAwsProviderRepoPath, resourceTypes)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> filename implementing resource")
	}

	resourceServices, err := terraform.GenerateResourceServiceMap(terraformAwsProviderRepoPath, outputPath,
		resourceTypes, resourceFileNames)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> AWS service")
	}

	resourceIDs, err := terraform.GenerateResourceIDMap(terraformAwsProviderRepoPath, outputPath, resourceFileNames)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> resource ID")
	}

	apis, err := aws.APIs(awsSdkGoRepoPath)
	if err != nil {
		log.WithError(err).Fatal("failed to load AWS APIs")
	}

	log.Infof("AWS services covered by Terraform: %d/%d",
		len(aws.ServicesCoveredByTerraform(resourceServices)), len(aws.ServicePkgNames(apis)))

	log.Debugf("AWS services not covered:")
	diff := util.Difference(aws.ServicePkgNames(apis), aws.ServicesCoveredByTerraform(resourceServices))
	for _, d := range diff {
		log.Debugf("\t%s", d)
	}

	log.Warn("AWS services used by Terraform that are named differently in AWS API v2:")
	diff = util.Difference(aws.ServicesCoveredByTerraform(resourceServices), aws.ServicePkgNames(apis))
	for _, d := range diff {
		log.Warnf("\t: %s", d)
	}

	err = aws.GenerateClient("../aws", aws.ServicesCoveredByTerraform(resourceServices))
	if err != nil {
		log.WithError(err).Fatal("Failed to write AWS client")
	}

	listFunctionNames, genResourceInfos := aws.GenerateListFunctions("../aws",
		resourceServices, resourceIDs, apis)

	log.Infof("Generated list functions: %d", len(listFunctionNames))

	err = terraform.GenerateListResourcesFunctions(outputPath, listFunctionNames)
	if err != nil {
		log.WithError(err).Fatal("failed to generate list resource functions (all and by type)")
	}

	err = aws.WriteReadme(outputPath, genResourceInfos)
	if err != nil {
		log.WithError(err).Fatal("failed to generate README")
	}
}
