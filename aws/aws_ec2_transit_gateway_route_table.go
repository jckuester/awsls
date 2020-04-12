// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGatewayRouteTable(client *Client) error {
	req := client.ec2conn.DescribeTransitGatewayRouteTablesRequest(&ec2.DescribeTransitGatewayRouteTablesInput{})

	p := ec2.NewDescribeTransitGatewayRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TransitGatewayRouteTables {
			fmt.Println(*r.TransitGatewayRouteTableId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.CreationTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
