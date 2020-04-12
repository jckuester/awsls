// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func ListCloudwatchLogGroup(client *Client) {
	req := client.cloudwatchlogsconn.DescribeLogGroupsRequest(&cloudwatchlogs.DescribeLogGroupsInput{})

	p := cloudwatchlogs.NewDescribeLogGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_cloudwatch_log_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LogGroups {
			fmt.Println(*r.LogGroupName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_cloudwatch_log_group: %s", err)
	}

}
