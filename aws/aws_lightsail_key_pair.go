// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailKeyPair(client *Client) {
	req := client.lightsailconn.GetKeyPairsRequest(&lightsail.GetKeyPairsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_lightsail_key_pair: %s", err)
	} else {
		if len(resp.KeyPairs) > 0 {
			fmt.Println("")
			fmt.Println("aws_lightsail_key_pair:")
			for _, r := range resp.KeyPairs {
				fmt.Println(*r.Name)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
