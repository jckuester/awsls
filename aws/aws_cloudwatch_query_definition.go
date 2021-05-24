// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudwatchQueryDefinition(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Cloudwatchlogsconn.DescribeQueryDefinitions(ctx, &cloudwatchlogs.DescribeQueryDefinitionsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.QueryDefinitions) > 0 {

		for _, r := range resp.QueryDefinitions {

			result = append(result, terraform.Resource{
				Type:      "aws_cloudwatch_query_definition",
				ID:        *r.QueryDefinitionId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
