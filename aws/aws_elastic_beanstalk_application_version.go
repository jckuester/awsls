// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElasticBeanstalkApplicationVersion(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Elasticbeanstalkconn.DescribeApplicationVersions(ctx, &elasticbeanstalk.DescribeApplicationVersionsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ApplicationVersions) > 0 {

		for _, r := range resp.ApplicationVersions {

			result = append(result, terraform.Resource{
				Type:      "aws_elastic_beanstalk_application_version",
				ID:        *r.ApplicationName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
