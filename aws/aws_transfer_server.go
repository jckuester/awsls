// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/transfer"
)

func ListTransferServer(client *Client) {
	req := client.transferconn.ListServersRequest(&transfer.ListServersInput{})

	p := transfer.NewListServersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_transfer_server:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Servers {
			fmt.Println(*r.ServerId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_transfer_server: %s", err)
	}

}
