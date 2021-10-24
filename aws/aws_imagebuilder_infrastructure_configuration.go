// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListImagebuilderInfrastructureConfiguration(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Imagebuilderconn.ListInfrastructureConfigurations(ctx, &imagebuilder.ListInfrastructureConfigurationsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.InfrastructureConfigurationSummaryList) > 0 {

		for _, r := range resp.InfrastructureConfigurationSummaryList {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, terraform.Resource{
				Type:      "aws_imagebuilder_infrastructure_configuration",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	return result, nil
}
