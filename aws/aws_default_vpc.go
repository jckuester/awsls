// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListDefaultVpc(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})

	var result []Resource

	p := ec2.NewDescribeVpcsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Vpcs {
			if *r.OwnerId != client.accountid {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type: "aws_default_vpc",
				ID:   *r.VpcId,
				Tags: tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
