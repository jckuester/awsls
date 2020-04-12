// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerEndpoint(client *Client) error {
	req := client.sagemakerconn.ListEndpointsRequest(&sagemaker.ListEndpointsInput{})

	p := sagemaker.NewListEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Endpoints {
			fmt.Println(*r.EndpointName)

			fmt.Printf("CreatedAt: %s\n", *r.CreationTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
