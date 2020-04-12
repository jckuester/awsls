// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafRuleGroup(client *Client) {
	req := client.wafconn.ListRuleGroupsRequest(&waf.ListRuleGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_waf_rule_group: %s", err)
	} else {
		if len(resp.RuleGroups) > 0 {
			fmt.Println("")
			fmt.Println("aws_waf_rule_group:")
			for _, r := range resp.RuleGroups {
				fmt.Println(*r.RuleGroupId)

			}
		}
	}

}
