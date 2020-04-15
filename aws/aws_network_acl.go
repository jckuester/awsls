// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNetworkAcl(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})

	var result []Resource

	p := ec2.NewDescribeNetworkAclsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NetworkAcls {
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type: "aws_network_acl",
				ID:   *r.NetworkAclId,
				Tags: tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
