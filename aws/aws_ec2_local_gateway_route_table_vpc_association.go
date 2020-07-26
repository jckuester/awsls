// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2LocalGatewayRouteTableVpcAssociation(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeLocalGatewayRouteTableVpcAssociationsRequest(&ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput{})

	var result []Resource

	p := ec2.NewDescribeLocalGatewayRouteTableVpcAssociationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LocalGatewayRouteTableVpcAssociations {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:    "aws_ec2_local_gateway_route_table_vpc_association",
				ID:      *r.LocalGatewayRouteTableVpcAssociationId,
				Profile: client.Profile,
				Region:  client.Region,
				Tags:    tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
