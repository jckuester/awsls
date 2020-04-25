// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNatGateway(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeNatGatewaysRequest(&ec2.DescribeNatGatewaysInput{})

	var result []Resource

	p := ec2.NewDescribeNatGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NatGateways {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreateTime
			result = append(result, Resource{
				Type:      "aws_nat_gateway",
				ID:        *r.NatGatewayId,
				Region:    client.ec2conn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
