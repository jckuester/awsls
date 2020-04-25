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

	"github.com/jckuester/awsls/gen/util"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func GenerateResourceIDMap(providerRepoPath, outputPath string, resourceFileNames map[string]string) (map[string]string, error) {
	resourceIDs := map[string]string{}
	for rType, fileName := range resourceFileNames {
		resourceID, err := GetResourceID(providerRepoPath, fileName)
		if err != nil {
			continue
		}
		resourceIDs[rType] = resourceID
	}

	err := writeResourceIDs(outputPath, resourceIDs)
	if err != nil {
		return nil, err
	}

	log.WithField("length", len(resourceIDs)).Infof("Generated map of resource type -> resource ID")

	return resourceIDs, nil
}
func GetResourceID(providerRepoPath, fileName string) (string, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset,
		fmt.Sprintf("%s/aws/%s", providerRepoPath, fileName),
		nil, 0)
	if err != nil {
		return "", err
	}

	var result *string

	ast.Inspect(node, func(n ast.Node) bool {
		// find d.SetId(...) function calls
		fn, ok := n.(*ast.CallExpr)
		if ok {
			if fun, ok := fn.Fun.(*ast.SelectorExpr); ok {
				funcName := fun.Sel.Name
				if funcName == "SetId" {
					// setId has only one argument
					switch x := fn.Args[0].(type) {

					case *ast.BasicLit:
					// ignore d.SetId("")
					case *ast.Ident:
						if strings.Contains(strings.ToLower(x.Name), "name") {
							// if the identifier contains "name", we know it's the
							// name identifying a resource but figure out the real AWS struct field name
							// at a later stage
							result = aws.String("NAME_PLACEHOLDER")
						}

						// handle the following kind of expressions:
						//   id := *res.ImageId
						//   d.SetId(id)
						ast.Inspect(node, func(n ast.Node) bool {
							ass, ok := n.(*ast.AssignStmt)
							if !ok {
								return true
							}

							if ass.Tok != token.DEFINE {
								return true
							}

							leftAss, ok := ass.Lhs[0].(ast.Expr)
							if !ok {
								return true
							}

							ident, ok := leftAss.(*ast.Ident)
							if !ok {
								return true
							}

							if ident.Name != x.Name {
								return true
							}

							rightAss, ok := ass.Rhs[0].(ast.Expr)
							if !ok {
								return true
							}

							switch x := rightAss.(type) {
							case *ast.StarExpr:
								identRight, ok := x.X.(*ast.SelectorExpr)
								if !ok {
									return true
								}

								result = &identRight.Sel.Name
							}

							return true
						})
					case *ast.StarExpr:
						// handle the following kind of expressions:
						//   d.SetId(*vpc.VpcId)
						ident, ok := x.X.(*ast.SelectorExpr)
						if !ok {
							return true
						}

						result = &ident.Sel.Name

					case *ast.TypeAssertExpr:
						// handle the following kind of expressions:
						//   d.SetId(d.Get("name").(string))
						call, ok := x.X.(*ast.CallExpr)
						if !ok {
							return true
						}

						fn, ok := call.Fun.(*ast.SelectorExpr)
						if !ok {
							return true
						}

						if fn.Sel.Name == "Get" {
							ident, ok := call.Args[0].(*ast.BasicLit)
							if !ok {
								return true
							}

							if strings.Contains(strings.ToLower(ident.Value), "name") {
								result = aws.String("NAME_PLACEHOLDER")
								return true
							}
						}

					case *ast.CallExpr:
						// handle the following kind of expressions:
						//   d.SetId(aws.StringValue(output.CertificateAuthorityArn))
						fn, ok := x.Fun.(*ast.SelectorExpr)
						if !ok {
							return true
						}

						if fn.Sel.Name == "StringValue" {
							selExpr, ok := x.Args[0].(*ast.SelectorExpr)
							if !ok {
								return true
							}

							result = &selExpr.Sel.Name

							return true
						}
					}
				}
			}
		}

		return true
	})

	if result != nil {
		return *result, nil
	}

	return "", fmt.Errorf("no ID found for resource type")
}

func writeResourceIDs(outputPath string, resourceIDs map[string]string) error {
	err := os.MkdirAll(outputPath, 0775)
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", err)
	}

	err = util.WriteGoFile(
		filepath.Join(outputPath, "ids.go"),
		util.CodeLayout,
		"",
		"resource",
		resourceIDsGoCode(resourceIDs),
	)
	if err != nil {
		return fmt.Errorf("failed to write Go code to file: %s", err)
	}

	return nil
}

func resourceIDsGoCode(resourceIDs map[string]string) string {
	var buf bytes.Buffer
	err := resourceIDsTmpl.Execute(&buf, resourceIDs)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())

}

var resourceIDsTmpl = template.Must(template.New("resourceIDs").Parse(`
// IDs stores the name of the struct field of the AWS API used as ID for each Terraform resource type.
var IDs = map[string]string{
{{range $key, $value := .}}"{{$key}}": "{{$value}}",
{{end}}}
`))
