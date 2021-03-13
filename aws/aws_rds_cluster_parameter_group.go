// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRdsClusterParameterGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := rds.NewDescribeDBClusterParameterGroupsPaginator(client.Rdsconn, &rds.DescribeDBClusterParameterGroupsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

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

	return result, nil
}
