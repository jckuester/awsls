// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpoint(client *Client) {
	req := client.ec2conn.DescribeVpcEndpointsRequest(&ec2.DescribeVpcEndpointsInput{})

	p := ec2.NewDescribeVpcEndpointsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_vpc_endpoint:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.VpcEndpoints {
			fmt.Println(*r.VpcEndpointId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_vpc_endpoint: %s", err)
	}

}
