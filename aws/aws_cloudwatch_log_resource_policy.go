// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func ListCloudwatchLogResourcePolicy(client *Client) {
	req := client.cloudwatchlogsconn.DescribeResourcePoliciesRequest(&cloudwatchlogs.DescribeResourcePoliciesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_cloudwatch_log_resource_policy: %s", err)
	} else {
		if len(resp.ResourcePolicies) > 0 {
			fmt.Println("")
			fmt.Println("aws_cloudwatch_log_resource_policy:")
			for _, r := range resp.ResourcePolicies {
				fmt.Println(*r.PolicyName)

			}
		}
	}

}
