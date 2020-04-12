// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func ListAutoscalingGroup(client *Client) error {
	req := client.autoscalingconn.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{})

	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.AutoScalingGroups {
			fmt.Println(*r.AutoScalingGroupName)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.CreatedTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
