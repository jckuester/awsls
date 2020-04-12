// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func ListApiGatewayRestApi(client *Client) {
	req := client.apigatewayconn.GetRestApisRequest(&apigateway.GetRestApisInput{})

	p := apigateway.NewGetRestApisPaginator(req)
	fmt.Println("")
	fmt.Println("aws_api_gateway_rest_api:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Items {
			fmt.Println(*r.Id)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_api_gateway_rest_api: %s", err)
	}

}
