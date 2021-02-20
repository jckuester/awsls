// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafregionalRuleGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Wafregionalconn.ListRuleGroupsRequest(&wafregional.ListRuleGroupsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RuleGroups) > 0 {

		for _, r := range resp.RuleGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_wafregional_rule_group",
				ID:        *r.RuleGroupId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
