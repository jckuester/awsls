// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func ListCloudwatchLogDestination(client *Client) error {
	req := client.cloudwatchlogsconn.DescribeDestinationsRequest(&cloudwatchlogs.DescribeDestinationsInput{})

	p := cloudwatchlogs.NewDescribeDestinationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Destinations {
			fmt.Println(*r.DestinationName)

			t := time.Unix(0, *r.CreationTime*1000000).UTC()
			fmt.Printf("CreatedAt: %s\n", t)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
