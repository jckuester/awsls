// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAutoscalingGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Autoscalingconn.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{})

	var result []terraform.Resource

	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.AutoScalingGroups {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_autoscaling_group",
				ID:        *r.AutoScalingGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
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
