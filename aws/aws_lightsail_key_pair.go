// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLightsailKeyPair(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Lightsailconn.GetKeyPairs(ctx, &lightsail.GetKeyPairsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.KeyPairs) > 0 {

		for _, r := range resp.KeyPairs {

			result = append(result, terraform.Resource{
				Type:      "aws_lightsail_key_pair",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
