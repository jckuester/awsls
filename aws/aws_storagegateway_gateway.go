// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
)

func ListStoragegatewayGateway(client *Client) {
	req := client.storagegatewayconn.ListGatewaysRequest(&storagegateway.ListGatewaysInput{})

	p := storagegateway.NewListGatewaysPaginator(req)
	fmt.Println("")
	fmt.Println("aws_storagegateway_gateway:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Gateways {
			fmt.Println(*r.GatewayARN)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_storagegateway_gateway: %s", err)
	}

}
