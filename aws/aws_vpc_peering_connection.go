// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcPeeringConnection(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeVpcPeeringConnectionsRequest(&ec2.DescribeVpcPeeringConnectionsInput{})

	var result []Resource

	p := ec2.NewDescribeVpcPeeringConnectionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.VpcPeeringConnections {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:      "aws_vpc_peering_connection",
				ID:        *r.VpcPeeringConnectionId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
