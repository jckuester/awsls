// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLambdaEventSourceMapping(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Lambdaconn.ListEventSourceMappingsRequest(&lambda.ListEventSourceMappingsInput{})

	var result []terraform.Resource

	p := lambda.NewListEventSourceMappingsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

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

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
