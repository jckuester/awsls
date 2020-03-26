package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/apex/log/handlers/cli"

	"github.com/apex/log"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	resourceTypes := ResourceTypes()
	log.Infof("Number of Terraform AWS resource types: %d", len(resourceTypes))

	for _, r := range resourceTypes {
		fileName := GetResourceTypeFileName(r)
		if fileName != nil {
			log.Debugf("%s: %s", r, *fileName)
		} else {
			log.Debugf(r)
		}
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

// GetResourceTypeFileName returns the name of the file that implements the resource of the given resource type.
func GetResourceTypeFileName(resourceType string) *string {
	node, err := parser.ParseFile(token.NewFileSet(),
		"/home/jan/git/github.com/yoyolabsio/terraform-provider-aws/aws/provider.go",
		nil, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var result *string

	ast.Inspect(node, func(n ast.Node) bool {
		mapEntry, ok := n.(*ast.KeyValueExpr)
		if !ok {
			return true
		}

		key, ok := mapEntry.Key.(*ast.BasicLit)
		if !ok {
			return true
		}

		if resourceType != strings.Trim(key.Value, "\"") {
			return false
		}

		resource, ok := mapEntry.Value.(*ast.CallExpr)
		if !ok {
			return false
		}

		resourceFuncName, ok := resource.Fun.(*ast.Ident)
		if !ok {
			return false
		}

		files, err := ioutil.ReadDir("/home/jan/git/github.com/yoyolabsio/terraform-provider-aws/aws")
		if err != nil {
			log.Fatal(err.Error())
		}

		for _, f := range files {
			if strings.HasPrefix(f.Name(), "resource_aws") {
				node, err := parser.ParseFile(token.NewFileSet(),
					"/home/jan/git/github.com/yoyolabsio/terraform-provider-aws/aws/"+f.Name(),
					nil, 0)
				if err != nil {
					log.Fatal(err.Error())
				}

				ast.Inspect(node, func(n ast.Node) bool {
					fn, ok := n.(*ast.FuncDecl)
					if !ok {
						return true
					}

					if fn.Name.Name == resourceFuncName.Name {
						result = aws.String(f.Name())
						return false
					}

					return true
				})
			}
		}
		return true
	})
	return result
}
