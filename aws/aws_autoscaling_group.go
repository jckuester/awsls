// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func ListAutoscalingGroup(client *Client) ([]Resource, error) {
	req := client.Autoscalingconn.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{})

	var result []Resource

	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.AutoScalingGroups {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreatedTime
			result = append(result, Resource{
				Type:      "aws_autoscaling_group",
				ID:        *r.AutoScalingGroupName,
				Region:    client.Autoscalingconn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
