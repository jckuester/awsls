// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEgressOnlyInternetGateway(client *Client) {
	req := client.ec2conn.DescribeEgressOnlyInternetGatewaysRequest(&ec2.DescribeEgressOnlyInternetGatewaysInput{})

	p := ec2.NewDescribeEgressOnlyInternetGatewaysPaginator(req)
	fmt.Println("")
	fmt.Println("aws_egress_only_internet_gateway:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EgressOnlyInternetGateways {
			fmt.Println(*r.EgressOnlyInternetGatewayId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_egress_only_internet_gateway: %s", err)
	}

}
