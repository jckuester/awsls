// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRuleGroup(client *Client) {
	req := client.wafregionalconn.ListRuleGroupsRequest(&wafregional.ListRuleGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_rule_group: %s", err)
	} else {
		if len(resp.RuleGroups) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_rule_group:")
			for _, r := range resp.RuleGroups {
				fmt.Println(*r.RuleGroupId)

			}
		}
	}

}
