// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListAlbListener(client *Client) {
	req := client.elasticloadbalancingv2conn.DescribeListenersRequest(&elasticloadbalancingv2.DescribeListenersInput{})

	p := elasticloadbalancingv2.NewDescribeListenersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_alb_listener:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Listeners {
			fmt.Println(*r.ListenerArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_alb_listener: %s", err)
	}

}
