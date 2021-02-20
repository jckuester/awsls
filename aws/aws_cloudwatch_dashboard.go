// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudwatchDashboard(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Cloudwatchconn.ListDashboardsRequest(&cloudwatch.ListDashboardsInput{})

	var result []terraform.Resource

	p := cloudwatch.NewListDashboardsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DashboardEntries {

			result = append(result, terraform.Resource{
				Type:      "aws_cloudwatch_dashboard",
				ID:        *r.DashboardName,
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
