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

//GenerateListResourcesByTypeFunction generates a function to list all resources of a given Terraform resource type.
func GenerateListResourcesByTypeFunction(outputPath string, listFunctionNames map[string]string) error {
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

terradozer "github.com/jckuester/terradozer/pkg/resource"
)

type Resource struct {
	Type string
	ID string
	Region string
    Profile string
	Tags map[string]string
	CreatedAt *time.Time
	terradozer.UpdatableResource
}

func ListResourcesByType(client *Client, resourceType string) ([]Resource, error) {
	switch resourceType {
	{{ range $key, $value := . }}case "{{ $key }}":
	return List{{ $value }}(client)
	{{ end }}default:
		return nil, fmt.Errorf("resource type is not (yet) supported: %s", resourceType)
	}
}
`))
