// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxLag(client *Client) error {
	req := client.directconnectconn.DescribeLagsRequest(&directconnect.DescribeLagsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Lags) > 0 {
		for _, r := range resp.Lags {
			fmt.Println(*r.LagId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
