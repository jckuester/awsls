// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListCustomerGateway(client *Client) error {
	req := client.ec2conn.DescribeCustomerGatewaysRequest(&ec2.DescribeCustomerGatewaysInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.CustomerGateways) > 0 {
		for _, r := range resp.CustomerGateways {
			fmt.Println(*r.CustomerGatewayId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
