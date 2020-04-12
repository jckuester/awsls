// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNetworkInterface(client *Client) error {
	req := client.ec2conn.DescribeNetworkInterfacesRequest(&ec2.DescribeNetworkInterfacesInput{})

	p := ec2.NewDescribeNetworkInterfacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NetworkInterfaces {
			fmt.Println(*r.NetworkInterfaceId)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
