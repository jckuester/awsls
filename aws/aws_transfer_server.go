// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/transfer"
)

func ListTransferServer(client *Client) error {
	req := client.transferconn.ListServersRequest(&transfer.ListServersInput{})

	p := transfer.NewListServersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Servers {
			fmt.Println(*r.ServerId)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
