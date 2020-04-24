// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2ClientVpnEndpoint(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeClientVpnEndpointsRequest(&ec2.DescribeClientVpnEndpointsInput{})

	var result []Resource

	p := ec2.NewDescribeClientVpnEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ClientVpnEndpoints {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.CreationTime)
			if err != nil {
				return nil, err
			}
			result = append(result, Resource{
				Type:      "aws_ec2_client_vpn_endpoint",
				ID:        *r.ClientVpnEndpointId,
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
