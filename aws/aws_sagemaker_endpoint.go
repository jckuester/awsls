// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerEndpoint(client *Client) {
	req := client.sagemakerconn.ListEndpointsRequest(&sagemaker.ListEndpointsInput{})

	p := sagemaker.NewListEndpointsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_sagemaker_endpoint:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Endpoints {
			fmt.Println(*r.EndpointName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_sagemaker_endpoint: %s", err)
	}

}
