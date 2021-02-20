// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListTransferServer(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Transferconn.ListServersRequest(&transfer.ListServersInput{})

	var result []terraform.Resource

	p := transfer.NewListServersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Servers {

			result = append(result, terraform.Resource{
				Type:      "aws_transfer_server",
				ID:        *r.ServerId,
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
