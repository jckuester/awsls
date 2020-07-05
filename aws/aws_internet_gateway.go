// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListInternetGateway(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeInternetGatewaysRequest(&ec2.DescribeInternetGatewaysInput{})

	var result []Resource

	p := ec2.NewDescribeInternetGatewaysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.InternetGateways {
			if *r.OwnerId != client.accountid {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_internet_gateway",
				ID:     *r.InternetGatewayId,
				Region: client.Ec2conn.Config.Region,
				Tags:   tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
