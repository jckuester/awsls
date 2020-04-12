// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListSubnet(client *Client) {
	req := client.ec2conn.DescribeSubnetsRequest(&ec2.DescribeSubnetsInput{})

	p := ec2.NewDescribeSubnetsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_subnet:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Subnets {
			fmt.Println(*r.SubnetId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_subnet: %s", err)
	}

}
