// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalRule(client *Client) {
	req := client.wafregionalconn.ListRulesRequest(&wafregional.ListRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_wafregional_rule: %s", err)
	} else {
		if len(resp.Rules) > 0 {
			fmt.Println("")
			fmt.Println("aws_wafregional_rule:")
			for _, r := range resp.Rules {
				fmt.Println(*r.RuleId)

			}
		}
	}

}
