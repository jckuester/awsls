// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbInstance(client *Client) error {
	req := client.rdsconn.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})

	p := rds.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBInstances {
			fmt.Println(*r.DBInstanceIdentifier)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
