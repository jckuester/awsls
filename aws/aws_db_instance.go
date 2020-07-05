// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbInstance(client *Client) ([]Resource, error) {
	req := client.Rdsconn.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})

	var result []Resource

	p := rds.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DBInstances {

			t := *r.InstanceCreateTime
			result = append(result, Resource{
				Type:   "aws_db_instance",
				ID:     *r.DBInstanceIdentifier,
				Region: client.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
