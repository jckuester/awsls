// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcPeeringConnection(client *Client) {
	req := client.ec2conn.DescribeVpcPeeringConnectionsRequest(&ec2.DescribeVpcPeeringConnectionsInput{})

	p := ec2.NewDescribeVpcPeeringConnectionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_vpc_peering_connection:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.VpcPeeringConnections {
			fmt.Println(*r.VpcPeeringConnectionId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_vpc_peering_connection: %s", err)
	}

}
