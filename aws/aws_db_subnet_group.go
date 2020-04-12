// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSubnetGroup(client *Client) error {
	req := client.rdsconn.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{})

	p := rds.NewDescribeDBSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBSubnetGroups {
			fmt.Println(*r.DBSubnetGroupName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
