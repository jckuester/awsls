// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbParameterGroup(client *Client) error {
	req := client.rdsconn.DescribeDBParameterGroupsRequest(&rds.DescribeDBParameterGroupsInput{})

	p := rds.NewDescribeDBParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBParameterGroups {
			fmt.Println(*r.DBParameterGroupName)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
