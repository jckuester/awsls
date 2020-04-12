// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func ListLambdaEventSourceMapping(client *Client) {
	req := client.lambdaconn.ListEventSourceMappingsRequest(&lambda.ListEventSourceMappingsInput{})

	p := lambda.NewListEventSourceMappingsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_lambda_event_source_mapping:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSourceMappings {
			fmt.Println(*r.UUID)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_lambda_event_source_mapping: %s", err)
	}

}
