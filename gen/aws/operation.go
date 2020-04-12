// +build codegen

package aws

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/private/model/api"
)

type Operation struct {
	api.Operation

	TerraformType  string
	ResourceID     string
	OutputListName string
	OpName         string
	GetTagsGoCode  string
}

func (o *Operation) GoCode() string {
	var buf bytes.Buffer
	err := listResourcesOperationTmpl.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var listResourcesOperationTmpl = template.Must(template.New("listResourcesOperation").Parse(`
import(
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/{{ .API.PackageName }}"
)

{{ $reqType := printf "%sRequest" .ExportedName -}}
{{ $respType := printf "%sResponse" .ExportedName -}}
{{ $pagerType := printf "%sPaginator" .ExportedName -}}

func  List{{.OpName}}(client *Client) {
    req := client.{{ .API.PackageName }}conn.{{ $reqType }}(&{{ .API.PackageName }}.{{ .InputRef.GoTypeElem }}{})

	{{ if .Paginator }}

    p := {{ .API.PackageName }}.New{{ $pagerType }}(req)
	fmt.Println("")
	fmt.Println("{{.TerraformType}}:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.{{ .OutputListName }}{
			fmt.Println(*r.{{ .ResourceID }})
			{{ if ne .GetTagsGoCode "" }}{{ .GetTagsGoCode }}{{ end }}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("{{ .TerraformType }}: %s", err)
	}

	{{ else }}

    resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("{{ .TerraformType }}: %s", err)
	} else {
		if len(resp.{{ .OutputListName }}) > 0 {
			fmt.Println("")
			fmt.Println("{{.TerraformType}}:")
			for _, r := range resp.{{ .OutputListName }}{
				fmt.Println(*r.{{ .ResourceID }})
				{{ if ne .GetTagsGoCode "" }}{{ .GetTagsGoCode }}{{ end }}
			}
		}
	}

	{{ end }}
}
`))
