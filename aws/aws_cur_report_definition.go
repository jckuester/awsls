// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCurReportDefinition(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Costandusagereportserviceconn.DescribeReportDefinitionsRequest(&costandusagereportservice.DescribeReportDefinitionsInput{})

	var result []terraform.Resource

	p := costandusagereportservice.NewDescribeReportDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ReportDefinitions {

			result = append(result, terraform.Resource{
				Type:      "aws_cur_report_definition",
				ID:        *r.ReportName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
