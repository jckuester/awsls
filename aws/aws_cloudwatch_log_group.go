// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func ListCloudwatchLogGroup(client *Client) ([]Resource, error) {
	req := client.cloudwatchlogsconn.DescribeLogGroupsRequest(&cloudwatchlogs.DescribeLogGroupsInput{})

	var result []Resource

	p := cloudwatchlogs.NewDescribeLogGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LogGroups {

			t := time.Unix(0, *r.CreationTime*1000000).UTC()
			result = append(result, Resource{
				Type:   "aws_cloudwatch_log_group",
				ID:     *r.LogGroupName,
				Region: client.cloudwatchlogsconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
