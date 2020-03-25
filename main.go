package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	resourceTypes := ResourceTypes()
	log.Infof("Number of Terraform AWS resource types: %d", len(resourceTypes))
	for _, r := range resourceTypes {
		log.Debugf(r)
	}
}

// ResourceTypes returns a list of all resource types supported by the Terraform AWS Provider.
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
