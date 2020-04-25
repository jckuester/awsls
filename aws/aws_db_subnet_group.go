// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSubnetGroup(client *Client) ([]Resource, error) {
	req := client.rdsconn.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{})

	var result []Resource

	p := rds.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBSubnetGroups {

			result = append(result, Resource{
				Type:   "aws_db_subnet_group",
				ID:     *r.DBSubnetGroupName,
				Region: client.rdsconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
