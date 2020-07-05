// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRuleGroup(client *Client) ([]Resource, error) {
	req := client.Wafregionalconn.ListRuleGroupsRequest(&wafregional.ListRuleGroupsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RuleGroups) > 0 {
		for _, r := range resp.RuleGroups {

			result = append(result, Resource{
				Type:   "aws_wafregional_rule_group",
				ID:     *r.RuleGroupId,
				Region: client.Wafregionalconn.Config.Region,
			})
		}
	}

	return result, nil
}
