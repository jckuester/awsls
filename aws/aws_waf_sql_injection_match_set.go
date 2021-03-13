// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafSqlInjectionMatchSet(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Wafconn.ListSqlInjectionMatchSets(ctx, &waf.ListSqlInjectionMatchSetsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.SqlInjectionMatchSets) > 0 {

		for _, r := range resp.SqlInjectionMatchSets {

			result = append(result, terraform.Resource{
				Type:      "aws_waf_sql_injection_match_set",
				ID:        *r.SqlInjectionMatchSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
