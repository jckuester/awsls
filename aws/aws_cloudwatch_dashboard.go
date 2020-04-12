// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func ListCloudwatchDashboard(client *Client) error {
	req := client.cloudwatchconn.ListDashboardsRequest(&cloudwatch.ListDashboardsInput{})

	p := cloudwatch.NewListDashboardsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DashboardEntries {
			fmt.Println(*r.DashboardName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
