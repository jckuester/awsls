// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

func ListElb(client *Client) ([]Resource, error) {
	req := client.elasticloadbalancingconn.DescribeLoadBalancersRequest(&elasticloadbalancing.DescribeLoadBalancersInput{})

	var result []Resource

	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LoadBalancerDescriptions {

			t := *r.CreatedTime
			result = append(result, Resource{
				Type:   "aws_elb",
				ID:     *r.LoadBalancerName,
				Region: client.elasticloadbalancingconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
