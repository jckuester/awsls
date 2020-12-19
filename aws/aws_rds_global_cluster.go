// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListRdsGlobalCluster(client *Client) ([]Resource, error) {
	req := client.Rdsconn.DescribeGlobalClustersRequest(&rds.DescribeGlobalClustersInput{})

	var result []Resource

	p := rds.NewDescribeGlobalClustersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.GlobalClusters {

			result = append(result, Resource{
				Type:      "aws_rds_global_cluster",
				ID:        *r.GlobalClusterIdentifier,
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
