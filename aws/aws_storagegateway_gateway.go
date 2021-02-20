// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListStoragegatewayGateway(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Storagegatewayconn.ListGatewaysRequest(&storagegateway.ListGatewaysInput{})

	var result []terraform.Resource

	p := storagegateway.NewListGatewaysPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Gateways {

			result = append(result, terraform.Resource{
				Type:      "aws_storagegateway_gateway",
				ID:        *r.GatewayARN,
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
