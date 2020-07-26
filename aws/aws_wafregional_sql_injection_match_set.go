// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalSqlInjectionMatchSet(client *Client) ([]Resource, error) {
	req := client.Wafregionalconn.ListSqlInjectionMatchSetsRequest(&wafregional.ListSqlInjectionMatchSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SqlInjectionMatchSets) > 0 {
		for _, r := range resp.SqlInjectionMatchSets {

			result = append(result, Resource{
				Type:    "aws_wafregional_sql_injection_match_set",
				ID:      *r.SqlInjectionMatchSetId,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	return result, nil
}
