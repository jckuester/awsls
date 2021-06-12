// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListConfigConfigurationAggregator(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Configserviceconn.DescribeConfigurationAggregators(ctx, &configservice.DescribeConfigurationAggregatorsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ConfigurationAggregators) > 0 {

		for _, r := range resp.ConfigurationAggregators {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_config_configuration_aggregator",
				ID:        *r.ConfigurationAggregatorName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
