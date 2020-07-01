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

## Supported Resource Types

Currently, all {{ .SupportedResourceTypeCount }} resource types across {{ len .Services }} services in the table below can be
listed with awsls. The "Tags" column shows if a resource supports displaying tags, the "Creation Time" column if a
resource has a creation timestamp, and the "Owner" column if resources are pre-filtered by account ID.

| Service / Type | Tags | Creation Time | Owner
| :------------- | :--: | :-----------: | :---:
{{ range .Services }}{{ $service := . }}| **{{ $service }}** |
{{ range $key, $value := $infos }}{{ if eq $key  $service -}}
{{ range $value -}}
| {{ .Type }} | {{ if .Tags }} x {{ end }} | {{ if .CreationTime }} x {{ end }} |{{ if .Owner }} x |{{ end }}
{{ end }}
{{- end }}{{ end }}{{ end }}
`))
