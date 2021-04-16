// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafregionalSqlInjectionMatchSet(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Wafregionalconn.ListSqlInjectionMatchSets(ctx, &wafregional.ListSqlInjectionMatchSetsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.SqlInjectionMatchSets) > 0 {

		for _, r := range resp.SqlInjectionMatchSets {

			result = append(result, terraform.Resource{
				Type:      "aws_wafregional_sql_injection_match_set",
				ID:        *r.SqlInjectionMatchSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
