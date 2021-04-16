// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafregionalXssMatchSet(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Wafregionalconn.ListXssMatchSets(ctx, &wafregional.ListXssMatchSetsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.XssMatchSets) > 0 {

		for _, r := range resp.XssMatchSets {

			result = append(result, terraform.Resource{
				Type:      "aws_wafregional_xss_match_set",
				ID:        *r.XssMatchSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
