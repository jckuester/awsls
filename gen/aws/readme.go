// +build codegen

package aws

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
)

func WriteReadme(outputPath string, resourceTypes []GeneratedResourceInfo) error {
	err := ioutil.WriteFile(filepath.Join(outputPath, "README.md"),
		[]byte(readmeCode(resourceTypes)), 0664)

	if err != nil {
		return fmt.Errorf("failed to write README.md: %s", err)
	}

	return nil
}

func readmeCode(terraformTypes []GeneratedResourceInfo) string {
	var buf bytes.Buffer
	err := Readme.Execute(&buf, terraformTypes)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var Readme = template.Must(template.New("readme").Parse(`

| Resource Type                    | Tag 
| :-----------------------------   |:-------------:
{{ range . }}| {{ .Type }}         |   {{ if .Tags }} x {{ end }} | 
{{ end }} 
`))
