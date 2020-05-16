// +build codegen

package aws

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type readmeData struct {
	Services                   []string
	ResourceInfos              map[string][]GeneratedResourceInfo
	SupportedResourceTypeCount int
}

func WriteReadme(outputPath string, resourceInfos map[string][]GeneratedResourceInfo) error {
	err := ioutil.WriteFile(filepath.Join(outputPath, "README.md"),
		[]byte(readmeCode(resourceInfos)), 0664)

	if err != nil {
		return fmt.Errorf("failed to write README.md: %s", err)
	}

	return nil
}

func readmeCode(resourceInfos map[string][]GeneratedResourceInfo) string {
	var buf bytes.Buffer

	resourceCount := 0

	var services []string
	for service, resources := range resourceInfos {
		resourceCount += len(resources)
		services = append(services, service)
	}

	sort.Strings(services)

	err := Readme.Execute(&buf, readmeData{
		Services:                   services,
		ResourceInfos:              resourceInfos,
		SupportedResourceTypeCount: resourceCount,
	})
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var Readme = template.Must(template.New("readme").Parse(`
{{ $infos := .ResourceInfos }}

# awsls

**Note: This tool is still WIP**

A list command for AWS resources. Supports listing of {{ .SupportedResourceTypeCount }} resource types across
{{ len .Services }} different services.


## How are so many resource types covered?

The answer is that awsls is mainly code generated; here is the
[code of the generator](./gen). Feel free to fork it and generate something else.

## Usage

	awsls <resource_type glob pattern>

To list all VPCs, for example, run

    awsls aws_vpc

or to list all resources

    awsls "*"

## Supported Resource Types

Currently, all resource types in the table below can be listed with awsls. The Tags, Creation Time, and Owner
column shows which resource type supports tags, has a creation date, or can be filtered by account owner, respectively.

| Service / Type | Tags | Creation Time | Owner
| :------------- | :--: | :-----------: | :---:
{{ range .Services }}{{ $service := . }}| **{{ $service }}** |
{{ range $key, $value := $infos }}{{ if eq $key  $service -}}
{{ range $value -}}
| {{ .Type }} | {{ if .Tags }} x {{ end }} | {{ if .CreationTime }} x {{ end }} |{{ if .Owner }} x |{{ end }}
{{ end }}
{{- end }}{{ end }}{{ end }}
`))
