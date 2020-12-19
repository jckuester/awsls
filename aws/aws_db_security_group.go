// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSecurityGroup(client *Client) ([]Resource, error) {
	req := client.Rdsconn.DescribeDBSecurityGroupsRequest(&rds.DescribeDBSecurityGroupsInput{})

	var result []Resource

	p := rds.NewDescribeDBSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DBSecurityGroups {
			if *r.OwnerId != client.AccountID {
				continue
			}

			result = append(result, Resource{
				Type:      "aws_db_security_group",
				ID:        *r.DBSecurityGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
