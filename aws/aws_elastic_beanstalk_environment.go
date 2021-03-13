// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElasticBeanstalkEnvironment(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Elasticbeanstalkconn.DescribeEnvironments(ctx, &elasticbeanstalk.DescribeEnvironmentsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.Environments) > 0 {

		for _, r := range resp.Environments {

			result = append(result, terraform.Resource{
				Type:      "aws_elastic_beanstalk_environment",
				ID:        *r.EnvironmentId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
