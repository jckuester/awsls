// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListAlbTargetGroup(client *Client) ([]Resource, error) {
	req := client.Elasticloadbalancingv2conn.DescribeTargetGroupsRequest(&elasticloadbalancingv2.DescribeTargetGroupsInput{})

	var result []Resource

	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TargetGroups {

			result = append(result, Resource{
				Type:   "aws_alb_target_group",
				ID:     *r.TargetGroupArn,
				Region: client.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
