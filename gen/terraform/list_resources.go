// +build codegen

package terraform

import (
	"bytes"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jckuester/awsls/gen/aws"

	"github.com/jckuester/awsls/gen/util"
)

//GenerateListResourcesByTypeFunction generates a function to list all resources of a given Terraform resource type.
func GenerateListResourcesByTypeFunction(outputPath string, rTypes []aws.ResourceType) {
	err := util.WriteGoFile(
		filepath.Join(outputPath, "list.go"),
		util.CodeLayout,
		"",
		"aws",
		listByTypeGoCode(rTypes),
	)

	if err != nil {
		panic(err)
	}
}

func listByTypeGoCode(rTypes []aws.ResourceType) string {
	var buf bytes.Buffer
	err := listByTypeTmpl.Execute(&buf, rTypes)
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
	AccountID string
	Tags map[string]string
	CreatedAt *time.Time
	terradozer.UpdatableResource
}

func ListResourcesByType(client *Client, resourceType string) ([]Resource, error) {
	switch resourceType {
	{{ range . }}case "{{ .Name }}":
	return {{ .ListFunctionName }}(client)
	{{ end }}default:
		return nil, fmt.Errorf("resource type is not (yet) supported: %s", resourceType)
	}
}
`))
