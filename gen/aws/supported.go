// +build codegen

package aws

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jckuester/awsls/gen/util"
)

// GenerateSupportedResourceTypeList generates code of a list of Terraform resource types that are currently
// supported by awsls and writes the code to directory outputPath.
func GenerateSupportedResourceTypeList(outputPath string, listFunctionNames map[string]string) error {
	err := util.WriteGoFile(
		filepath.Join(outputPath, "supported.go"),
		util.CodeLayout,
		"",
		"resource",
		supportedResourcesGoCode(listFunctionNames),
	)

	if err != nil {
		return fmt.Errorf("failed to write Go code to file: %s", err)
	}

	return nil
}

func supportedResourcesGoCode(listFunctionNames map[string]string) string {
	var buf bytes.Buffer
	err := supportedResourcesTmpl.Execute(&buf, listFunctionNames)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var supportedResourcesTmpl = template.Must(template.New("supportedResources").Parse(`
// SupportedTypes is a list of all resource types currently supported by awsls.
var SupportedTypes = []string{
{{ range $key, $value := . }}"{{ $key }}",
{{ end }}}
`))
