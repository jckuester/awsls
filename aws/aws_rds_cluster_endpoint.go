// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRdsClusterEndpoint(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Rdsconn.DescribeDBClusterEndpointsRequest(&rds.DescribeDBClusterEndpointsInput{})

	var result []terraform.Resource

	p := rds.NewDescribeDBClusterEndpointsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

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

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
