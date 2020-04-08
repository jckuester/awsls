// +build codegen

package util

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/private/util"
)

const CodeLayout = `// Code is generated. DO NOT EDIT.

%s
package %s

%s
`

func WriteGoFile(file string, layout string, args ...interface{}) error {
	return ioutil.WriteFile(file, []byte(util.GoFmt(fmt.Sprintf(layout, args...))), 0664)
}

func Difference(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 1; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
