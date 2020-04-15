// +build codegen

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
		"aws",
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

var listByTypeTmpl = template.Must(template.New("listByType").Parse(`import(
"fmt"
"time"
)

type Resource struct {
	Type string
	ID string
	Tags map[string]string
	CreatedAt *time.Time
}

func ListResources(client *Client) error {
{{ range $key, $value := . }}List{{ $value }}(client)
{{end}}
return nil
}

func ListResourcesByType(client *Client, resourceType string) ([]Resource, error) {
	switch resourceType {
	{{range $key, $value := .}}case "{{ $key }}":
	return List{{ $value }}(client)
	{{end}}default:
		return nil, fmt.Errorf("resource type is not (yet) supported: %s", resourceType)
	}
}
`))
