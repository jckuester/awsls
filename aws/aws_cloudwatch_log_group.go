// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func ListCloudwatchLogGroup(client *Client) error {
	req := client.cloudwatchlogsconn.DescribeLogGroupsRequest(&cloudwatchlogs.DescribeLogGroupsInput{})

	p := cloudwatchlogs.NewDescribeLogGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LogGroups {
			fmt.Println(*r.LogGroupName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
