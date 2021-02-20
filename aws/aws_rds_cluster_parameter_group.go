// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRdsClusterParameterGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Rdsconn.DescribeDBClusterParameterGroupsRequest(&rds.DescribeDBClusterParameterGroupsInput{})

	var result []terraform.Resource

	p := rds.NewDescribeDBClusterParameterGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DBClusterParameterGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_rds_cluster_parameter_group",
				ID:        *r.DBClusterParameterGroupName,
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
