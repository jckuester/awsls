// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGateway(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeTransitGatewaysRequest(&ec2.DescribeTransitGatewaysInput{})

	var result []Resource

	p := ec2.NewDescribeTransitGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TransitGateways {
			if *r.OwnerId != client.accountid {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_ec2_transit_gateway",
				ID:        *r.TransitGatewayId,
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
