// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpointService(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeVpcEndpointServicesRequest(&ec2.DescribeVpcEndpointServicesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ServiceDetails) > 0 {
		for _, r := range resp.ServiceDetails {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_vpc_endpoint_service",
				ID:     *r.ServiceId,
				Region: client.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
