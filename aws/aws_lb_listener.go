// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListLbListener(client *Client) error {
	req := client.elasticloadbalancingv2conn.DescribeListenersRequest(&elasticloadbalancingv2.DescribeListenersInput{})

	p := elasticloadbalancingv2.NewDescribeListenersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Listeners {
			fmt.Println(*r.ListenerArn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
