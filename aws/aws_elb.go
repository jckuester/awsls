// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

func ListElb(client *Client) error {
	req := client.elasticloadbalancingconn.DescribeLoadBalancersRequest(&elasticloadbalancing.DescribeLoadBalancersInput{})

	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LoadBalancerDescriptions {
			fmt.Println(*r.LoadBalancerName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
