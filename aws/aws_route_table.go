// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListRouteTable(client *Client) {
	req := client.ec2conn.DescribeRouteTablesRequest(&ec2.DescribeRouteTablesInput{})

	p := ec2.NewDescribeRouteTablesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_route_table:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.RouteTables {
			fmt.Println(*r.RouteTableId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_route_table: %s", err)
	}

}
