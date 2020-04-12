// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpnGateway(client *Client) error {
	req := client.ec2conn.DescribeVpnGatewaysRequest(&ec2.DescribeVpnGatewaysInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.VpnGateways) > 0 {
		for _, r := range resp.VpnGateways {
			fmt.Println(*r.VpnGatewayId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
