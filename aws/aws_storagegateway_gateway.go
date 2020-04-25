// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
)

func ListStoragegatewayGateway(client *Client) ([]Resource, error) {
	req := client.storagegatewayconn.ListGatewaysRequest(&storagegateway.ListGatewaysInput{})

	var result []Resource

	p := storagegateway.NewListGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Gateways {

			result = append(result, Resource{
				Type:   "aws_storagegateway_gateway",
				ID:     *r.GatewayARN,
				Region: client.storagegatewayconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
