// +build codegen

package aws

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/private/model/api"
	"github.com/jckuester/awsls/gen/util"
)

type Service struct {
	Name                   string
	TerraformResourceTypes []ResourceType
}

type ResourceType struct {
	Name         string
	Tags         bool
	CreationTime bool
	Owner        bool
}

func GenerateListFunctions(outputPath string, services []Service, resourceIDs map[string]string,
	resourceTypesWithTags []string, apis api.APIs) []Service {

	resourcesWithRequiredFieldsCount := 0
	noOutputFieldNameFoundCount := 0
	noListOpCandidatesFoundCount := 0
	noResourceIDFoundCount := 0

	var servicesResult []Service

	for _, service := range services {
		fmt.Println()
		fmt.Printf("service: %s\n---\n", service.Name)

		var rTypesResult []ResourceType

		for _, rType := range service.TerraformResourceTypes {
			_, ok := ExcludedResourceTypes[rType.Name]
			if ok {
				log.WithField("resource", rType).Info("exclude")
				continue
			}

			listOpCandidates := FindListOperationCandidates(rType.Name, service.Name, apis)
			if len(listOpCandidates) == 0 {
				noListOpCandidatesFoundCount++
				log.WithField("resource", rType).Errorf("no list operation candidate found")

				continue
			}

			outputFieldName, op, err := findOutputField(rType.Name, listOpCandidates, "structure")
			if err != nil {
				_, _, err = findOutputField(rType.Name, listOpCandidates, "string")
				if err != nil {
					noOutputFieldNameFoundCount++
					log.WithError(err).WithField("resource", rType.Name).Errorf("unable to find output field name")

					continue
				}

				log.WithField("resource", rType.Name).Infof("found output field of type string")

				continue
			}

			outputField := op.OutputRef.Shape.MemberRefs[outputFieldName]

			if len(op.InputRef.Shape.Required) > 0 {
				resourcesWithRequiredFieldsCount++
				log.WithField("resource", rType).Errorf("required input fields: %s", op.InputRef.Shape.Required)

				continue
			}

			resourceID, err := findResourceID(rType.Name, resourceIDs, outputField)
			if err != nil {
				noResourceIDFoundCount++
				log.WithField("resource", rType).Errorf("no resource ID found")

				continue
			}

			for k, _ := range op.InputRef.Shape.MemberRefs {
				if strings.Contains(strings.ToLower(k), "owner") {
					log.Infof("input; found owner field for %s: %s", rType, k)
				}
			}

			op.OutputFieldName = outputFieldName
			op.TerraformType = rType.Name
			op.ResourceID = resourceID
			op.OpName = rType.ListFunctionName()
			op.Inputs = Inputs[rType.Name]

			rType.CreationTime = op.GetCreationTimeGoCode() != ""
			rType.Owner = op.GetOwnerGoCode() != ""
			rType.Tags = util.Contains(resourceTypesWithTags, rType.Name)

			if rType.Name != "aws_instance" {
				// note: code is manually added for "aws_instance"
				writeListFunction(outputPath, &op)
			} else {
				rType.CreationTime = true
			}

			rTypesResult = append(rTypesResult, rType)
		}

		if len(rTypesResult) > 0 {
			servicesResult = append(servicesResult, Service{
				Name:                   service.Name,
				TerraformResourceTypes: rTypesResult,
			})
		}
	}

	log.Infof("list functions with required fields: %d", resourcesWithRequiredFieldsCount)
	log.Infof("unable to find output field name: %d", noOutputFieldNameFoundCount)
	log.Infof("resources without list operation candidate: %d", noListOpCandidatesFoundCount)
	log.Infof("no resource ID found: %d", noResourceIDFoundCount)

	return servicesResult
}

func writeListFunction(outputPath string, op *ListOperation) {
	err := util.WriteGoFile(
		filepath.Join(outputPath, op.TerraformType+".go"),
		util.CodeLayout,
		"",
		"aws",
		op.GoCode(),
	)

	if err != nil {
		panic(err)
	}
}

type ListOperation struct {
	api.Operation

	TerraformType   string
	ResourceID      string
	OutputListName  string
	OpName          string
	Inputs          string
	OutputFieldName string
}

func (o *ListOperation) GoCode() string {
	var buf bytes.Buffer
	err := listResourcesOperationTmpl.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

func (o ListOperation) GetTagsGoCode() string {
	outputField := o.OutputRef.Shape.MemberRefs[o.OutputFieldName]

	for k, v := range outputField.Shape.MemberRef.Shape.MemberRefs {
		if k == "Tags" {
			if v.Shape.Type == "list" {
				return `tags := map[string]string{}
						for _, t := range r.Tags {
							tags[*t.Key] = *t.Value
						}`
			}

			if v.Shape.Type == "map" {
				return `tags := map[string]string{}
						for k, v := range r.Tags {
							tags[k] = v
						}`
			}
		}

		if strings.Contains(k, "Tag") {
			log.Infof("tags: %s %s", k, v.Shape.Type)
		}
	}

	return ""
}

func (o ListOperation) GetCreationTimeGoCode() string {
	outputField := o.OutputRef.Shape.MemberRefs[o.OutputFieldName]

	creationTimeFieldNames := []string{
		"LaunchTime",
		"CreateTime",
		"CreateDate",
		"CreatedTime",
		"CreationDate",
		"CreationTime",
		"CreationTimestamp",
		"StartTime",
		"InstanceCreateTime",
	}

	for k, v := range outputField.Shape.MemberRef.Shape.MemberRefs {
		for _, name := range creationTimeFieldNames {
			if k == name {
				if v.Shape.Type == "string" {
					return `t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.` + k + `)
							if err != nil {
								return nil, err
							}`
				}

				if v.Shape.Type == "timestamp" {
					return `t := ` + fmt.Sprintf("*r.%s", k)
				}

				if v.Shape.Type == "long" {
					return fmt.Sprintf("t := time.Unix(0, *r.%s * 1000000).UTC()", k)
				}

				log.Warnf("uncovered creation time type: %s", v.Shape.Type)
			}
		}
	}

	return ""
}

func (o ListOperation) GetOwnerGoCode() string {
	outputField := o.OutputRef.Shape.MemberRefs[o.OutputFieldName]

	for k, _ := range outputField.Shape.MemberRef.Shape.MemberRefs {
		if k == "OwnerId" {
			return `if *r.OwnerId != client.AccountID {
						continue
					}`
		}
		if strings.Contains(strings.ToLower(k), "owner") {
			log.Infof("output; found owner field: %s", k)
		}
	}

	return ""
}

var listResourcesOperationTmpl = template.Must(template.New("listResourcesOperation").Funcs(
	template.FuncMap{
		"Title": strings.Title,
	}).Parse(`
import(
	"context"
	"github.com/aws/aws-sdk-go-v2/service/{{ .API.PackageName }}"
)

{{ $reqType := printf "%sRequest" .ExportedName -}}
{{ $pagerType := printf "%sPaginator" .ExportedName -}}

func  {{.OpName}}(client *Client) ([]Resource, error) {
    req := client.{{ .API.PackageName | Title }}conn.{{ $reqType }}(&{{ .API.PackageName }}.{{ .InputRef.GoTypeElem }}{ {{ if ne .Inputs "" }}{{ .Inputs }}{{ end }} })

	var result []Resource

	{{ if .Paginator }}
    p := {{ .API.PackageName }}.New{{ $pagerType }}(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()
	{{ else }}
    resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.{{ .OutputListName }}) > 0 {
	{{ end }}
		for _, r := range resp.{{ .OutputListName }}{
			{{ if ne .GetOwnerGoCode "" }}{{ .GetOwnerGoCode }}{{ end }}
			{{ if ne .GetTagsGoCode "" }}{{ .GetTagsGoCode }}{{ end }}
			{{ if ne .GetCreationTimeGoCode "" }}{{ .GetCreationTimeGoCode }}{{ end }}
			result = append(result, Resource{
				Type: "{{ .TerraformType }}",
				ID: *r.{{ .ResourceID }},
				Profile: client.Profile,
				Region: client.Region,
				AccountID: client.AccountID,
				{{ if ne .GetTagsGoCode "" }}Tags: tags,{{ end }}
				{{ if ne .GetCreationTimeGoCode "" }}CreatedAt: &t,{{ end }}
			})
		}
	}

	{{ if .Paginator }}
	if err := p.Err(); err != nil {
		return nil, err
	}
	{{ end }}

	return result, nil
}
`))
