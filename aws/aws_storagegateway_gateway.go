// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
)

func ListStoragegatewayGateway(client *Client) error {
	req := client.storagegatewayconn.ListGatewaysRequest(&storagegateway.ListGatewaysInput{})

	p := storagegateway.NewListGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Gateways {
			fmt.Println(*r.GatewayARN)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
