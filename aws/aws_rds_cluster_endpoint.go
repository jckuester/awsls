// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRdsClusterEndpoint(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := rds.NewDescribeDBClusterEndpointsPaginator(client.Rdsconn, &rds.DescribeDBClusterEndpointsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.DBClusterEndpoints {

			result = append(result, terraform.Resource{
				Type:      "aws_rds_cluster_endpoint",
				ID:        *r.DBClusterEndpointIdentifier,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
