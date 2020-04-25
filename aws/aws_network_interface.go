// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNetworkInterface(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeNetworkInterfacesRequest(&ec2.DescribeNetworkInterfacesInput{})

	var result []Resource

	p := ec2.NewDescribeNetworkInterfacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NetworkInterfaces {
			if *r.OwnerId != client.accountid {
				continue
			}

			result = append(result, Resource{
				Type:   "aws_network_interface",
				ID:     *r.NetworkInterfaceId,
				Region: client.ec2conn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
