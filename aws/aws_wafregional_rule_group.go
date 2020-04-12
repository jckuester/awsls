// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRuleGroup(client *Client) error {
	req := client.wafregionalconn.ListRuleGroupsRequest(&wafregional.ListRuleGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.RuleGroups) > 0 {
		for _, r := range resp.RuleGroups {
			fmt.Println(*r.RuleGroupId)
		}
	}

	return nil
}
