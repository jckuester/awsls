// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafSizeConstraintSet(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Wafconn.ListSizeConstraintSets(ctx, &waf.ListSizeConstraintSetsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.SizeConstraintSets) > 0 {

		for _, r := range resp.SizeConstraintSets {

			result = append(result, terraform.Resource{
				Type:      "aws_waf_size_constraint_set",
				ID:        *r.SizeConstraintSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
