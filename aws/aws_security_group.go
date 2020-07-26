// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListSecurityGroup(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeSecurityGroupsRequest(&ec2.DescribeSecurityGroupsInput{})

	var result []Resource

	p := ec2.NewDescribeSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecurityGroups {
			if *r.OwnerId != client.AccountID {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:    "aws_security_group",
				ID:      *r.GroupId,
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
