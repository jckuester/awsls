// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TransitGatewayRouteTable(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeTransitGatewayRouteTablesRequest(&ec2.DescribeTransitGatewayRouteTablesInput{})

	var result []Resource

	p := ec2.NewDescribeTransitGatewayRouteTablesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.TransitGatewayRouteTables {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_ec2_transit_gateway_route_table",
				ID:        *r.TransitGatewayRouteTableId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
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
