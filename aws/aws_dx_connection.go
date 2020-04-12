// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxConnection(client *Client) error {
	req := client.directconnectconn.DescribeConnectionsRequest(&directconnect.DescribeConnectionsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Connections) > 0 {
		for _, r := range resp.Connections {
			fmt.Println(*r.ConnectionId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
