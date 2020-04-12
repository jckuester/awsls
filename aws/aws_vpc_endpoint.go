// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpoint(client *Client) error {
	req := client.ec2conn.DescribeVpcEndpointsRequest(&ec2.DescribeVpcEndpointsInput{})

	p := ec2.NewDescribeVpcEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.VpcEndpoints {
			fmt.Println(*r.VpcEndpointId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.CreationTimestamp)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
