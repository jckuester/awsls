// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEgressOnlyInternetGateway(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeEgressOnlyInternetGatewaysRequest(&ec2.DescribeEgressOnlyInternetGatewaysInput{})

	var result []Resource

	p := ec2.NewDescribeEgressOnlyInternetGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EgressOnlyInternetGateways {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type: "aws_egress_only_internet_gateway",
				ID:   *r.EgressOnlyInternetGatewayId,
				Tags: tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
