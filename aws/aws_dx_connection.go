// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxConnection(client *Client) {
	req := client.directconnectconn.DescribeConnectionsRequest(&directconnect.DescribeConnectionsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_dx_connection: %s", err)
	} else {
		if len(resp.Connections) > 0 {
			fmt.Println("")
			fmt.Println("aws_dx_connection:")
			for _, r := range resp.Connections {
				fmt.Println(*r.ConnectionId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
