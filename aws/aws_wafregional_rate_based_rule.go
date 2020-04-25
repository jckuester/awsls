// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRateBasedRule(client *Client) ([]Resource, error) {
	req := client.wafregionalconn.ListRateBasedRulesRequest(&wafregional.ListRateBasedRulesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) > 0 {
		for _, r := range resp.Rules {

			result = append(result, Resource{
				Type:   "aws_wafregional_rate_based_rule",
				ID:     *r.RuleId,
				Region: client.wafregionalconn.Config.Region,
			})
		}
	}

	return result, nil
}
