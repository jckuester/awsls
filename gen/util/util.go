//go:build codegen

package util

import (
	"fmt"
	"go/format"
	"io/ioutil"
)

const CodeLayout = `// Code is generated. DO NOT EDIT.

%s
package %s

%s
`

func WriteGoFile(file string, layout string, args ...interface{}) error {
	return ioutil.WriteFile(file, []byte(GoFmt(fmt.Sprintf(layout, args...))), 0664)
}

// GoFmt returns the Go formatted string of the input.
//
// Panics if the format fails.
func GoFmt(buf string) string {
	formatted, err := format.Source([]byte(buf))
	if err != nil {
		panic(fmt.Errorf("%s\nOriginal code:\n%s", err.Error(), buf))
	}

	return string(formatted)
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

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
