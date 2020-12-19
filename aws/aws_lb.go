// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListLb(client *Client) ([]Resource, error) {
	req := client.Elasticloadbalancingv2conn.DescribeLoadBalancersRequest(&elasticloadbalancingv2.DescribeLoadBalancersInput{})

	var result []Resource

	p := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.LoadBalancers {

			t := *r.CreatedTime
			result = append(result, Resource{
				Type:      "aws_lb",
				ID:        *r.LoadBalancerArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
