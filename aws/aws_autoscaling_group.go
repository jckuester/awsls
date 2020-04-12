// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func ListAutoscalingGroup(client *Client) {
	req := client.autoscalingconn.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{})

	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_autoscaling_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.AutoScalingGroups {
			fmt.Println(*r.AutoScalingGroupName)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_autoscaling_group: %s", err)
	}

}
