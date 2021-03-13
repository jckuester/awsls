// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafRegexMatchSet(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Wafconn.ListRegexMatchSets(ctx, &waf.ListRegexMatchSetsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.RegexMatchSets) > 0 {

		for _, r := range resp.RegexMatchSets {

			result = append(result, terraform.Resource{
				Type:      "aws_waf_regex_match_set",
				ID:        *r.RegexMatchSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
