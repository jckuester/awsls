// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListKeyPair(client *Client) {
	req := client.ec2conn.DescribeKeyPairsRequest(&ec2.DescribeKeyPairsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_key_pair: %s", err)
	} else {
		if len(resp.KeyPairs) > 0 {
			fmt.Println("")
			fmt.Println("aws_key_pair:")
			for _, r := range resp.KeyPairs {
				fmt.Println(*r.KeyName)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
