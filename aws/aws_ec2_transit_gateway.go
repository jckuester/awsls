// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGateway(client *Client) {
	req := client.ec2conn.DescribeTransitGatewaysRequest(&ec2.DescribeTransitGatewaysInput{})

	p := ec2.NewDescribeTransitGatewaysPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ec2_transit_gateway:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TransitGateways {
			fmt.Println(*r.TransitGatewayId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ec2_transit_gateway: %s", err)
	}

}
