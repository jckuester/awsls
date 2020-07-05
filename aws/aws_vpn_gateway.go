// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpnGateway(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeVpnGatewaysRequest(&ec2.DescribeVpnGatewaysInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.VpnGateways) > 0 {
		for _, r := range resp.VpnGateways {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_vpn_gateway",
				ID:     *r.VpnGatewayId,
				Region: client.Ec2conn.Config.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
