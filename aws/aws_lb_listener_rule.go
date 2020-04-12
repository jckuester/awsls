// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListLbListenerRule(client *Client) {
	req := client.elasticloadbalancingv2conn.DescribeRulesRequest(&elasticloadbalancingv2.DescribeRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_lb_listener_rule: %s", err)
	} else {
		if len(resp.Rules) > 0 {
			fmt.Println("")
			fmt.Println("aws_lb_listener_rule:")
			for _, r := range resp.Rules {
				fmt.Println(*r.RuleArn)

			}
		}
	}

}
