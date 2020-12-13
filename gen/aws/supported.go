// +build codegen

package aws

import (
	"bytes"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jckuester/awsls/gen/util"
)

// GenerateListOfSupportedResourceTypes code-generates a list of Terraform resource types
// that are currently supported by awsls.
func GenerateListOfSupportedResourceTypes(outputPath string, rTypes []ResourceType) {
	err := util.WriteGoFile(
		filepath.Join(outputPath, "supported.go"),
		util.CodeLayout,
		"",
		"resource",
		supportedResourcesGoCode(rTypes),
	)

	if err != nil {
		panic(err)
	}
}

func supportedResourcesGoCode(rTypes []ResourceType) string {
	var buf bytes.Buffer
	err := supportedResourcesTmpl.Execute(&buf, rTypes)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var supportedResourcesTmpl = template.Must(template.New("supportedResources").Parse(`
// SupportedTypes is a list of all resource types currently supported by awsls.
var SupportedTypes = []string{
  {{- range . }}
    "{{ .Name }}",
  {{- end }}
}
`))
