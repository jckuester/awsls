// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func ListCloudwatchDashboard(client *Client) ([]Resource, error) {
	req := client.cloudwatchconn.ListDashboardsRequest(&cloudwatch.ListDashboardsInput{})

	var result []Resource

	p := cloudwatch.NewListDashboardsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DashboardEntries {

			result = append(result, Resource{
				Type: "aws_cloudwatch_dashboard",
				ID:   *r.DashboardName,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
