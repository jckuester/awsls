// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSecurityGroup(client *Client) error {
	req := client.rdsconn.DescribeDBSecurityGroupsRequest(&rds.DescribeDBSecurityGroupsInput{})

	p := rds.NewDescribeDBSecurityGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBSecurityGroups {
			fmt.Println(*r.DBSecurityGroupName)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
