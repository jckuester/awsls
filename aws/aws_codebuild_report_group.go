// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodebuildReportGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codebuildconn.ListReportGroupsRequest(&codebuild.ListReportGroupsInput{})

	var result []terraform.Resource

	p := codebuild.NewListReportGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ReportGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_codebuild_report_group",
				ID:        r,
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
