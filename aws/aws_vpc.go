// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpc(client *Client) error {
	req := client.ec2conn.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})

	p := ec2.NewDescribeVpcsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Vpcs {
			fmt.Println(*r.VpcId)
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
