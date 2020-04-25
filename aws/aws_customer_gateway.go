// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListCustomerGateway(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeCustomerGatewaysRequest(&ec2.DescribeCustomerGatewaysInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.CustomerGateways) > 0 {
		for _, r := range resp.CustomerGateways {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_customer_gateway",
				ID:     *r.CustomerGatewayId,
				Region: client.ec2conn.Config.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
