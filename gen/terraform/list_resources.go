package terraform

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jckuester/awsls/gen/util"
)

//GenerateListResourcesFunctions generates two functions: one to list all resources and one to list resources by type.
func GenerateListResourcesFunctions(outputPath string, listFunctionNames map[string]string) error {
	err := util.WriteGoFile(
		filepath.Join(outputPath, "list.go"),
		util.CodeLayout,
		"",
		"resource",
		listByTypeGoCode(listFunctionNames),
	)

	if err != nil {
		return fmt.Errorf("failed to write Go code to file: %s", err)
	}

	return nil
}

func listByTypeGoCode(listFunctionNames map[string]string) string {
	var buf bytes.Buffer
	err := listByTypeTmpl.Execute(&buf, listFunctionNames)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var listByTypeTmpl = template.Must(template.New("listByType").Parse(`import "github.com/jckuester/awsls/aws"

func ListResources(client *aws.Client) {
{{ range $key, $value := . }}aws.List{{ $value }}(client)
{{end}}}

func ListResourcesByType(client *aws.Client, resourceType string) {
	switch resourceType {
	{{range $key, $value := .}}case "{{ $key }}":
	aws.List{{ $value }}(client)
	{{end}}}
}
`))
