// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListConfigConformancePack(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Configserviceconn.DescribeConformancePacks(ctx, &configservice.DescribeConformancePacksInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ConformancePackDetails) > 0 {

		for _, r := range resp.ConformancePackDetails {

			result = append(result, terraform.Resource{
				Type:      "aws_config_conformance_pack",
				ID:        *r.ConformancePackName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
