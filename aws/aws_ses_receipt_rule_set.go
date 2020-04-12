// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func ListSesReceiptRuleSet(client *Client) {
	req := client.sesconn.ListReceiptRuleSetsRequest(&ses.ListReceiptRuleSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ses_receipt_rule_set: %s", err)
	} else {
		if len(resp.RuleSets) > 0 {
			fmt.Println("")
			fmt.Println("aws_ses_receipt_rule_set:")
			for _, r := range resp.RuleSets {
				fmt.Println(*r.Name)

			}
		}
	}

}
