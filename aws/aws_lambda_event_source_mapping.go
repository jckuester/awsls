// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func ListLambdaEventSourceMapping(client *Client) error {
	req := client.lambdaconn.ListEventSourceMappingsRequest(&lambda.ListEventSourceMappingsInput{})

	p := lambda.NewListEventSourceMappingsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSourceMappings {
			fmt.Println(*r.UUID)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
