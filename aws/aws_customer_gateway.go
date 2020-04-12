// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListCustomerGateway(client *Client) {
	req := client.ec2conn.DescribeCustomerGatewaysRequest(&ec2.DescribeCustomerGatewaysInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_customer_gateway: %s", err)
	} else {
		if len(resp.CustomerGateways) > 0 {
			fmt.Println("")
			fmt.Println("aws_customer_gateway:")
			for _, r := range resp.CustomerGateways {
				fmt.Println(*r.CustomerGatewayId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
