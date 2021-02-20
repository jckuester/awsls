// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDbInstance(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Rdsconn.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{})

	var result []terraform.Resource

	p := rds.NewDescribeDBInstancesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DBInstances {

			t := *r.InstanceCreateTime
			result = append(result, terraform.Resource{
				Type:      "aws_db_instance",
				ID:        *r.DBInstanceIdentifier,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
