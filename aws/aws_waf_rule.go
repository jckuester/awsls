// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafRule(client *Client) ([]Resource, error) {
	req := client.wafconn.ListRulesRequest(&waf.ListRulesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) > 0 {
		for _, r := range resp.Rules {

			result = append(result, Resource{
				Type:   "aws_waf_rule",
				ID:     *r.RuleId,
				Region: client.wafconn.Config.Region,
			})
		}
	}

	return result, nil
}
