// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func ListLambdaEventSourceMapping(client *Client) ([]Resource, error) {
	req := client.lambdaconn.ListEventSourceMappingsRequest(&lambda.ListEventSourceMappingsInput{})

	var result []Resource

	p := lambda.NewListEventSourceMappingsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSourceMappings {

			result = append(result, Resource{
				Type:   "aws_lambda_event_source_mapping",
				ID:     *r.UUID,
				Region: client.lambdaconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
