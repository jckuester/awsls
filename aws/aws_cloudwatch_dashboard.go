// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func ListCloudwatchDashboard(client *Client) {
	req := client.cloudwatchconn.ListDashboardsRequest(&cloudwatch.ListDashboardsInput{})

	p := cloudwatch.NewListDashboardsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_cloudwatch_dashboard:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DashboardEntries {
			fmt.Println(*r.DashboardName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_cloudwatch_dashboard: %s", err)
	}

}
