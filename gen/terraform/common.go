package terraform

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/private/util"
)

const codeLayout = `// Code is generated. DO NOT EDIT.

%s
package %s

%s
`

func writeGoFile(file string, layout string, args ...interface{}) error {
	return ioutil.WriteFile(file, []byte(util.GoFmt(fmt.Sprintf(layout, args...))), 0664)
}
