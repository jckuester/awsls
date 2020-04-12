// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListLbListenerRule(client *Client) error {
	req := client.elasticloadbalancingv2conn.DescribeRulesRequest(&elasticloadbalancingv2.DescribeRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Rules) > 0 {
		for _, r := range resp.Rules {
			fmt.Println(*r.RuleArn)

		}
	}

	return nil
}
