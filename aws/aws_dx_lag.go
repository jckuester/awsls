// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxLag(client *Client) {
	req := client.directconnectconn.DescribeLagsRequest(&directconnect.DescribeLagsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_dx_lag: %s", err)
	} else {
		if len(resp.Lags) > 0 {
			fmt.Println("")
			fmt.Println("aws_dx_lag:")
			for _, r := range resp.Lags {
				fmt.Println(*r.LagId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
