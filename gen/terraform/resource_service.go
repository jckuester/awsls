// +build codegen

package terraform

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/apex/log"

	"github.com/jckuester/awsls/gen/util"
)

func GenerateResourceServiceMap(providerRepoPath string, outputPath string, resourceTypes []string,
	resourceFileNames map[string]string) map[string]string {
	resourceServices := ResourceServices(providerRepoPath, resourceTypes, resourceFileNames)

	writeResourceServices(outputPath, resourceServices)

	log.WithField("length", len(resourceServices)).Infof("Generated map of resource type -> AWS service")

	return resourceServices
}

// resourceService returns the AWS service that the Terraform resource belongs to.
func ResourceServices(providerRepoPath string, resourceTypes []string,
	resourceFileNames map[string]string) map[string]string {
	resourceServices := map[string]string{}
	for _, rType := range resourceTypes {
		rFileName, ok := resourceFileNames[rType]
		if !ok {
			log.WithField("resource", rType).Warn("cannot find filename for resource type")
		}

		service, err := resourceService(providerRepoPath, rType, rFileName)
		if err != nil {
			log.WithError(err).WithField("resource", rType).Warn("could not identify AWS service for resource")
			continue
		}
		resourceServices[rType] = service
	}
	return resourceServices
}

func resourceService(providerRepoPath, resourceType, resourceFileName string) (string, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset,
		fmt.Sprintf("%s/aws/%s", providerRepoPath, resourceFileName),
		nil, 0)
	if err != nil {
		return "", err
	}

	var serviceCandidates []string

	for _, i := range node.Imports {
		importPath := strings.Trim(i.Path.Value, "\"")

		if strings.HasPrefix(importPath, "github.com/aws/aws-sdk-go/service/") {
			service := strings.TrimPrefix(importPath, "github.com/aws/aws-sdk-go/service/")
			serviceCandidates = append(serviceCandidates, service)
		}
	}

	if len(serviceCandidates) == 1 {
		service := serviceCandidates[0]

		return service, nil
	}

	if len(serviceCandidates) > 1 {
		manualService, ok := ManualResourceServiceMap[resourceType]
		if ok {
			return manualService, nil
		} else {
			return "", fmt.Errorf("multiple service candidates found: %s", serviceCandidates)
		}
	}

	return "", fmt.Errorf("no service candidate found")
}

func writeResourceServices(outputPath string, resourceServices map[string]string) {
	err := util.WriteGoFile(
		filepath.Join(outputPath, "services.go"),
		util.CodeLayout,
		"",
		"resource",
		resourceServicesGoCode(resourceServices),
	)

	if err != nil {
		panic(err)
	}
}

func resourceServicesGoCode(terraformTypes map[string]string) string {
	var buf bytes.Buffer
	err := resourceServicesTmpl.Execute(&buf, terraformTypes)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var resourceServicesTmpl = template.Must(template.New("resourceServices").Parse(`
// Services contains the name of the AWS service that each resource type belongs to.
var Services = map[string]string{
{{range $key, $value := .}}"{{$key}}": "{{$value}}",
{{end}}}
`))
