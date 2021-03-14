// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLightsailDomain(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Lightsailconn.GetDomains(ctx, &lightsail.GetDomainsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.Domains) > 0 {

		for _, r := range resp.Domains {

			result = append(result, terraform.Resource{
				Type:      "aws_lightsail_domain",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
