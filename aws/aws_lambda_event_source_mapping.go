// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLambdaEventSourceMapping(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := lambda.NewListEventSourceMappingsPaginator(client.Lambdaconn, &lambda.ListEventSourceMappingsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.EventSourceMappings {

			result = append(result, terraform.Resource{
				Type:      "aws_lambda_event_source_mapping",
				ID:        *r.UUID,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
