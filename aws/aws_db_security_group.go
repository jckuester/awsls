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
		page := p.CurrentPage()

		for _, r := range page.DBSecurityGroups {
			if *r.OwnerId != client.Accountid {
				continue
			}

			result = append(result, Resource{
				Type:   "aws_db_security_group",
				ID:     *r.DBSecurityGroupName,
				Region: client.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
