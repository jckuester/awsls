// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafRateBasedRule(client *Client) {
	req := client.wafconn.ListRateBasedRulesRequest(&waf.ListRateBasedRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_waf_rate_based_rule: %s", err)
	} else {
		if len(resp.Rules) > 0 {
			fmt.Println("")
			fmt.Println("aws_waf_rate_based_rule:")
			for _, r := range resp.Rules {
				fmt.Println(*r.RuleId)

			}
		}
	}

}
