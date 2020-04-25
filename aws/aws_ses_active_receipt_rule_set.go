// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func ListSesActiveReceiptRuleSet(client *Client) ([]Resource, error) {
	req := client.sesconn.ListReceiptRuleSetsRequest(&ses.ListReceiptRuleSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RuleSets) > 0 {
		for _, r := range resp.RuleSets {

			result = append(result, Resource{
				Type:   "aws_ses_active_receipt_rule_set",
				ID:     *r.Name,
				Region: client.sesconn.Config.Region,
			})
		}
	}

	return result, nil
}
