// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListLbListener(client *Client) ([]Resource, error) {
	req := client.elasticloadbalancingv2conn.DescribeListenersRequest(&elasticloadbalancingv2.DescribeListenersInput{})

	var result []Resource

	p := elasticloadbalancingv2.NewDescribeListenersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Listeners {

			result = append(result, Resource{
				Type: "aws_lb_listener",
				ID:   *r.ListenerArn,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
