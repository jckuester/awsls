// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafregionalIpset(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Wafregionalconn.ListIPSets(ctx, &wafregional.ListIPSetsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.IPSets) > 0 {

		for _, r := range resp.IPSets {

			result = append(result, terraform.Resource{
				Type:      "aws_wafregional_ipset",
				ID:        *r.IPSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
