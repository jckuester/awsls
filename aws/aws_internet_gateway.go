// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListInternetGateway(client *Client) {
	req := client.ec2conn.DescribeInternetGatewaysRequest(&ec2.DescribeInternetGatewaysInput{})

	p := ec2.NewDescribeInternetGatewaysPaginator(req)
	fmt.Println("")
	fmt.Println("aws_internet_gateway:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.InternetGateways {
			fmt.Println(*r.InternetGatewayId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_internet_gateway: %s", err)
	}

}
