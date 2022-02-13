//go:build codegen

package aws

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
)

type readmeData struct {
	Services                   []Service
	SupportedResourceTypeCount int
}

func GenerateReadme(outputPath string, services []Service, numOfResourceTypes int) {
	err := ioutil.WriteFile(filepath.Join(outputPath, "README.md"), []byte(readmeCode(services, numOfResourceTypes)), 0664)
	if err != nil {
		panic(err)
	}
}

func readmeCode(services []Service, numOfResourceTypes int) string {
	var buf bytes.Buffer

	tpl := template.Must(template.New("README.tpl.md").ParseFiles("../gen/aws/README.tpl.md"))

	err := tpl.Execute(&buf, readmeData{
		Services:                   services,
		SupportedResourceTypeCount: numOfResourceTypes,
	})

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String()) + "\n"
}
