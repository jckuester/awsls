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

	TerraformType         string
	ResourceID            string
	OutputListName        string
	OpName                string
	GetTagsGoCode         string
	GetCreationTimeGoCode string
	Imports               []string
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
	{{ range .Imports }}"{{ . }}"
	{{ end }}
	"github.com/aws/aws-sdk-go-v2/service/{{ .API.PackageName }}"
)

{{ $reqType := printf "%sRequest" .ExportedName -}}
{{ $respType := printf "%sResponse" .ExportedName -}}
{{ $pagerType := printf "%sPaginator" .ExportedName -}}

func  List{{.OpName}}(client *Client) ([]Resource, error) {
    req := client.{{ .API.PackageName }}conn.{{ $reqType }}(&{{ .API.PackageName }}.{{ .InputRef.GoTypeElem }}{})

	var result []Resource

	{{ if .Paginator }}
    p := {{ .API.PackageName }}.New{{ $pagerType }}(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.{{ .OutputListName }}{
			{{ if ne .GetTagsGoCode "" }}{{ .GetTagsGoCode }}{{ end }}
			{{ if ne .GetCreationTimeGoCode "" }}{{ .GetCreationTimeGoCode }}{{ end }}
			result = append(result, Resource{
				Type: "{{ .TerraformType }}",
				ID: *r.{{ .ResourceID }},
				{{ if ne .GetTagsGoCode "" }}Tags: tags,{{ end }}
				{{ if ne .GetCreationTimeGoCode "" }}CreatedAt: &t,{{ end }}
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	{{ else }}

    resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.{{ .OutputListName }}) > 0 {
		for _, r := range resp.{{ .OutputListName }}{
			{{ if ne .GetTagsGoCode "" }}{{ .GetTagsGoCode }}{{ end }}
			{{ if ne .GetCreationTimeGoCode "" }}{{ .GetCreationTimeGoCode }}{{ end }}
			result = append(result, Resource{
				Type: "{{ .TerraformType }}",
				ID: *r.{{ .ResourceID }},
				{{ if ne .GetTagsGoCode "" }}Tags: tags,{{ end }}
				{{ if ne .GetCreationTimeGoCode "" }}CreatedAt: &t,{{ end }}
			})
		}
	}
	{{ end }}
	return result, nil
}
`))
