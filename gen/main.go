// +build codegen

package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/aws/aws-sdk-go-v2/private/model/api"
	"github.com/jckuester/awsls/gen/aws"
	"github.com/jckuester/awsls/gen/terraform"
	"github.com/jckuester/awsls/gen/util"
)

const (
	outputPath       = "../resource"
	providerRepoPath = "/home/jan/git/github.com/terraform-provider-aws"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	resourceTypes, err := terraform.GenerateResourceTypeList(providerRepoPath, outputPath)
	if err != nil {
		log.WithError(err).Fatal("failed to generate list of Terraform AWS resource types")
	}

	resourceFileNames, err := terraform.ResourceFileNames(providerRepoPath, resourceTypes)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> filename implementing resource")
	}

	resourceServices, err := terraform.GenerateResourceServiceMap(providerRepoPath, outputPath,
		resourceTypes, resourceFileNames)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> AWS service")
	}

	resourceIDs, err := terraform.GenerateResourceIDMap(providerRepoPath, outputPath, resourceFileNames)
	if err != nil {
		log.WithError(err).Fatal("failed to generate map of resource type -> resource ID")
	}

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
		len(aws.ServicesCoveredByTerraform(resourceServices)), len(aws.Services(apis)))

	log.Debugf("AWS services not covered:")
	diff := util.Difference(aws.Services(apis), aws.ServicesCoveredByTerraform(resourceServices))
	for _, d := range diff {
		log.Debugf("\t%s", d)
	}

	diff = util.Difference(aws.ServicesCoveredByTerraform(resourceServices), aws.Services(apis))
	for _, d := range diff {
		log.Warnf("\tnot in v2 API: %s", d)
	}

	err = aws.WriteClient("../aws", aws.ServicesCoveredByTerraform(resourceServices))
	if err != nil {
		log.Fatal(err.Error())
	}
}
