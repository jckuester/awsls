// +build codegen

package terraform

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/jckuester/terratools/gen/util"
)

// ResourceTypes returns a list of all resource types implemented by the Terraform AWS Provider.
func ResourceTypes() []string {
	node, err := parser.ParseFile(token.NewFileSet(),
		"/home/jan/git/github.com/yoyolabsio/terraform-provider-aws/aws/provider.go",
		nil, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var result []string

	ast.Inspect(node, func(n ast.Node) bool {
		// look for a map called "ResourcesMap" that contains all the resource types
		resourcesMap, ok := n.(*ast.KeyValueExpr)
		if !ok {
			return true
		}

		resourceMapName, ok := resourcesMap.Key.(*ast.Ident)
		if !ok {
			return false
		}

		if resourceMapName.Name != "ResourcesMap" {
			return false
		}

		ast.Inspect(resourcesMap, func(n ast.Node) bool {
			mapEntry, ok := n.(*ast.KeyValueExpr)
			if !ok {
				return true
			}

			resourceType, ok := mapEntry.Key.(*ast.BasicLit)
			if !ok {
				return true
			}

			result = append(result, strings.Trim(resourceType.Value, "\""))

			return true
		})

		return true
	})

	return result
}

func WriteResourceTypes(outputPath string, resourceTypes []string) error {
	err := os.MkdirAll(outputPath, 0775)
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", err)
	}

	err = util.WriteGoFile(
		filepath.Join(outputPath, "resource_types.go"),
		util.CodeLayout,
		"",
		"resource",
		ResourceTypesGoCode(resourceTypes),
	)
	if err != nil {
		return fmt.Errorf("failed to write list of resource types to file: %s", err)
	}

	return nil
}

// ResourceTypesGoCode generates the Go code for the list of Terraform resource types.
func ResourceTypesGoCode(terraformTypes []string) string {
	var buf bytes.Buffer
	err := resourceTypesTmpl.Execute(&buf, terraformTypes)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var resourceTypesTmpl = template.Must(template.New("resourceTypes").Parse(`
// ResourceTypes is a list of all resource types implemented by the Terraform AWS Provider.
var ResourceTypes = []string{
{{range .}}"{{.}}",
{{end}}}
`))
