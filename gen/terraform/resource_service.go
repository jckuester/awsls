package terraform

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ManualResourceServiceMap are resources for which its corresponding service couldn't be automatically identified.
var ManualResourceServiceMap = map[string]string{
	"aws_alb":                                 "elasticloadbalancingv2",
	"aws_autoscaling_group":                   "autoscaling",
	"aws_efs_mount_target":                    "efs",
	"aws_elastic_beanstalk_environment":       "elasticbeanstalk",
	"aws_elb":                                 "elasticloadbalancing",
	"aws_iam_server_certificate":              "iam",
	"aws_lambda_event_source_mapping":         "lambda",
	"aws_launch_configuration":                "autoscaling",
	"aws_lb":                                  "elasticloadbalancingv2",
	"aws_s3_bucket_object":                    "s3",
	"aws_s3_bucket_public_access_block":       "s3",
	"aws_wafregional_byte_match_set":          "wafregional",
	"aws_wafregional_geo_match_set":           "wafregional",
	"aws_wafregional_ipset":                   "wafregional",
	"aws_wafregional_rate_based_rule":         "wafregional",
	"aws_wafregional_regex_match_set":         "wafregional",
	"aws_wafregional_regex_pattern_set":       "wafregional",
	"aws_wafregional_rule":                    "wafregional",
	"aws_wafregional_rule_group":              "wafregional",
	"aws_wafregional_size_constraint_set":     "wafregional",
	"aws_wafregional_sql_injection_match_set": "wafregional",
	"aws_wafregional_web_acl":                 "wafregional",
	"aws_wafregional_xss_match_set":           "wafregional",
}

// AWSServicesV1toV2 is a mapping from service names in AWS API v1 (used by Terraform)
// that have different names in v2 (used by this project in the generated code).
var AWSServicesV1toV2 = map[string]string{
	"elb":   "elasticloadbalancing",
	"elbv2": "elasticloadbalancingv2",
}

// ResourceService returns the AWS service that the Terraform resource belongs to.
func ResourceService(resourceType, resourceFileName string) (string, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset,
		"/home/jan/git/github.com/yoyolabsio/terraform-provider-aws/aws/"+resourceFileName,
		nil, 0)
	if err != nil {
		return "", err
	}

	var serviceCandidates []string

	for _, i := range node.Imports {
		importPath := strings.Trim(i.Path.Value, "\"")

		if strings.HasPrefix(importPath, "github.com/aws/aws-sdk-go/service/") {
			service := strings.TrimPrefix(importPath, "github.com/aws/aws-sdk-go/service/")
			serviceCandidates = append(serviceCandidates, service)
		}
	}

	if len(serviceCandidates) == 1 {
		service := serviceCandidates[0]

		serviceV2, ok := AWSServicesV1toV2[service]
		if ok {
			service = serviceV2
		}

		return service, nil
	}

	if len(serviceCandidates) > 1 {
		manualService, ok := ManualResourceServiceMap[resourceType]
		if ok {
			return manualService, nil
		} else {
			return "", fmt.Errorf("multiple service candidates found: %s", serviceCandidates)
		}
	}

	return "", fmt.Errorf("no service candidate found")
}

func WriteResourceServices(outputPath string, resourceServices map[string]string) error {
	err := os.MkdirAll(outputPath, 0775)
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", err)
	}

	err = writeGoFile(
		filepath.Join(outputPath, "resourceServices.go"),
		codeLayout,
		"",
		"resource",
		ResourceServicesGoCode(resourceServices),
	)
	if err != nil {
		return fmt.Errorf("failed to write services map to file: %s", err)
	}

	return nil
}

// ResourceServicesGoCode generates the Go code for the map of Terraform resource services.
func ResourceServicesGoCode(terraformTypes map[string]string) string {
	var buf bytes.Buffer
	err := resourceServicesTmpl.Execute(&buf, terraformTypes)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

var resourceServicesTmpl = template.Must(template.New("resourceServices").Parse(`
// ResourceServices contains the name of the AWS service that each resource type belongs to.
var ResourceServices = map[string]string{
{{range $key, $value := .}}"{{$key}}": "{{$value}}",
{{end}}}
`))
