// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
)

func ListWafSqlInjectionMatchSet(client *Client) ([]Resource, error) {
	req := client.Wafconn.ListSqlInjectionMatchSetsRequest(&waf.ListSqlInjectionMatchSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SqlInjectionMatchSets) > 0 {
		for _, r := range resp.SqlInjectionMatchSets {

			result = append(result, Resource{
				Type:   "aws_waf_sql_injection_match_set",
				ID:     *r.SqlInjectionMatchSetId,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
