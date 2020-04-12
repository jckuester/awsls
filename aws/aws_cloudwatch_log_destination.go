// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func ListCloudwatchLogDestination(client *Client) {
	req := client.cloudwatchlogsconn.DescribeDestinationsRequest(&cloudwatchlogs.DescribeDestinationsInput{})

	p := cloudwatchlogs.NewDescribeDestinationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_cloudwatch_log_destination:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Destinations {
			fmt.Println(*r.DestinationName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_cloudwatch_log_destination: %s", err)
	}

}
