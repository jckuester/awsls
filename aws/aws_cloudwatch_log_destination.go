// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func ListCloudwatchLogDestination(client *Client) ([]Resource, error) {
	req := client.Cloudwatchlogsconn.DescribeDestinationsRequest(&cloudwatchlogs.DescribeDestinationsInput{})

	var result []Resource

	p := cloudwatchlogs.NewDescribeDestinationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Destinations {

			t := time.Unix(0, *r.CreationTime*1000000).UTC()
			result = append(result, Resource{
				Type:   "aws_cloudwatch_log_destination",
				ID:     *r.DestinationName,
				Region: client.Cloudwatchlogsconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
