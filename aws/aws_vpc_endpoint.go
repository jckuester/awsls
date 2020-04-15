// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpoint(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeVpcEndpointsRequest(&ec2.DescribeVpcEndpointsInput{})

	var result []Resource

	p := ec2.NewDescribeVpcEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.VpcEndpoints {
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTimestamp
			result = append(result, Resource{
				Type:      "aws_vpc_endpoint",
				ID:        *r.VpcEndpointId,
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
