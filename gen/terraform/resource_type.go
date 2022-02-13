//go:build codegen

package terraform

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/jckuester/awsls/gen/util"
)

// GenerateResourceTypeList generates code of a list of Terraform resource types
// and writes the code to directory outputPath.
func GenerateResourceTypeList(providerRepoPath, outputPath string) ([]string, error) {
	resourceTypes, err := ResourceTypes(providerRepoPath)
	if err != nil {
		return nil, err
	}

	writeResourceTypes(outputPath, resourceTypes)

	log.WithField("length", len(resourceTypes)).Infof("Generated list of Terraform AWS resource types")
	return resourceTypes, nil
}

// ResourceTypes returns a list of all resource types implemented by the Terraform AWS Provider.
func ResourceTypes(providerRepoPath string) ([]string, error) {
	node, err := parser.ParseFile(token.NewFileSet(),
		fmt.Sprintf("%s/%s", providerRepoPath, "aws/provider.go"),
		nil, 0)
	if err != nil {
		return nil, err
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

	sort.Strings(result)

	return result, nil
}

func writeResourceTypes(outputPath string, resourceTypes []string) {
	code := resourceTypesGoCode(resourceTypes)

	err := util.WriteGoFile(
		filepath.Join(outputPath, "types.go"),
		util.CodeLayout,
		"",
		"resource",
		code,
	)

	if err != nil {
		panic(err)
	}
}

func resourceTypesGoCode(terraformTypes []string) string {
	var buf bytes.Buffer
	err := resourceTypesTmpl.Execute(&buf, terraformTypes)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var resourceTypesTmpl = template.Must(template.New("resourceTypes").Parse(`
// Types is a list of all resource types implemented by the Terraform AWS Provider.
var Types = []string{
{{range .}}"{{.}}",
{{end}}}
`))
