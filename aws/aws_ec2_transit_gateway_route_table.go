// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGatewayRouteTable(client *Client) {
	req := client.ec2conn.DescribeTransitGatewayRouteTablesRequest(&ec2.DescribeTransitGatewayRouteTablesInput{})

	p := ec2.NewDescribeTransitGatewayRouteTablesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ec2_transit_gateway_route_table:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TransitGatewayRouteTables {
			fmt.Println(*r.TransitGatewayRouteTableId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ec2_transit_gateway_route_table: %s", err)
	}

}
