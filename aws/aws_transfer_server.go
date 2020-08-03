// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/transfer"
)

func ListTransferServer(client *Client) ([]Resource, error) {
	req := client.Transferconn.ListServersRequest(&transfer.ListServersInput{})

	var result []Resource

	p := transfer.NewListServersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Servers {

			result = append(result, Resource{
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
