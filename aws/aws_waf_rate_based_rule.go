// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafRateBasedRule(client *Client) error {
	req := client.wafconn.ListRateBasedRulesRequest(&waf.ListRateBasedRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Rules) > 0 {
		for _, r := range resp.Rules {
			fmt.Println(*r.RuleId)

		}
	}

	return nil
}
