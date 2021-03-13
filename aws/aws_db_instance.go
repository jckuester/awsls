// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDbInstance(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := rds.NewDescribeDBInstancesPaginator(client.Rdsconn, &rds.DescribeDBInstancesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

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

	return result, nil
}
