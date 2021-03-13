// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCurReportDefinition(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := costandusagereportservice.NewDescribeReportDefinitionsPaginator(client.Costandusagereportserviceconn, &costandusagereportservice.DescribeReportDefinitionsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

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

	return result, nil
}
