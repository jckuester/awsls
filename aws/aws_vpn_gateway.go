// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpnGateway(client *Client) {
	req := client.ec2conn.DescribeVpnGatewaysRequest(&ec2.DescribeVpnGatewaysInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_vpn_gateway: %s", err)
	} else {
		if len(resp.VpnGateways) > 0 {
			fmt.Println("")
			fmt.Println("aws_vpn_gateway:")
			for _, r := range resp.VpnGateways {
				fmt.Println(*r.VpnGatewayId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
