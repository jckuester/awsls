// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEip(client *Client) error {
	req := client.ec2conn.DescribeAddressesRequest(&ec2.DescribeAddressesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Addresses) > 0 {
		for _, r := range resp.Addresses {
			fmt.Println(*r.AllocationId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
