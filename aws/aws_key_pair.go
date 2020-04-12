// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListKeyPair(client *Client) error {
	req := client.ec2conn.DescribeKeyPairsRequest(&ec2.DescribeKeyPairsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.KeyPairs) > 0 {
		for _, r := range resp.KeyPairs {
			fmt.Println(*r.KeyName)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
