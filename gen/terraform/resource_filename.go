//go:build codegen

package terraform

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws"
)

// ResourceFileNames returns for each Terraform resource type
// the name of the file in the Terraform AWS Provider code that implements it.
func ResourceFileNames(providerRepoPath string, resourceTypes []string) (map[string]string, error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", providerRepoPath, "aws"))
	if err != nil {
		return nil, err
	}

	funcDeclsPerFile := map[string][]string{}

	// store for all files implementing the resources their function declarations,
	// such as resourceAwsVpc().
	//
	// the idea of finding a file name is to link the information about a resource type in provider.go
	// (e.g., "aws_vpc": resourceAwsVpc()) with the declaration resourceAwsVpc() in the file aws/resource_aws_vpc.go.
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "resource_aws") &&
			!strings.HasSuffix(f.Name(), "_test.go") &&
			!strings.HasSuffix(f.Name(), "_migrate.go") {
			node, err := parser.ParseFile(token.NewFileSet(),
				fmt.Sprintf("%s/aws/%s", providerRepoPath, f.Name()),
				nil, 0)
			if err != nil {
				return nil, err
			}

			var funcDecls []string
			for _, decl := range node.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}

				funcDecls = append(funcDecls, fn.Name.Name)
			}
			funcDeclsPerFile[f.Name()] = funcDecls
		}
	}

	result := map[string]string{}

	for _, rType := range resourceTypes {
		resourceFileName, err := resourceFileName(providerRepoPath, rType, funcDeclsPerFile)
		if err != nil {
			log.WithField("resource", rType).Warn(err.Error())
			continue

		}
		result[rType] = resourceFileName
	}

	return result, nil
}

func resourceFileName(providerRepoPath, resourceType string, declarationsByFile map[string][]string) (string, error) {
	node, err := parser.ParseFile(token.NewFileSet(),
		fmt.Sprintf("%s/aws/provider.go", providerRepoPath),
		nil, 0)
	if err != nil {
		return "", err
	}

	var result *string

	ast.Inspect(node, func(n ast.Node) bool {
		// we are looking for map entries, such as "aws_vpc": resourceAwsVpc()
		// in the ResourcesMap in provider.go
		mapEntry, ok := n.(*ast.KeyValueExpr)
		if !ok {
			return true
		}

		key, ok := mapEntry.Key.(*ast.BasicLit)
		if !ok {
			return true
		}

		// e.g.: "aws_vpc"
		if resourceType != strings.Trim(key.Value, "\"") {
			return false
		}

		// e.g.:  resourceAwsVpc()
		callExpr, ok := mapEntry.Value.(*ast.CallExpr)
		if !ok {
			return false
		}

		resourceFuncName, ok := callExpr.Fun.(*ast.Ident)
		if !ok {
			return false
		}

		for fileName, fnNames := range declarationsByFile {
			for _, fnName := range fnNames {
				if fnName == resourceFuncName.Name {
					result = aws.String(fileName)
					return false
				}
			}
		}

		return true
	})

	if result != nil {
		return *result, nil
	}

	return "", fmt.Errorf("no file found that implements this resource type")
}
