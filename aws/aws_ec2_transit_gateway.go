// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGateway(client *Client) error {
	req := client.ec2conn.DescribeTransitGatewaysRequest(&ec2.DescribeTransitGatewaysInput{})

	p := ec2.NewDescribeTransitGatewaysPaginator(req)
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
		return err
	}

	return nil
}
