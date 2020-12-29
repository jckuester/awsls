// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListRdsCluster(client *Client) ([]Resource, error) {
	req := client.Rdsconn.DescribeDBClustersRequest(&rds.DescribeDBClustersInput{})

	var result []Resource

	p := rds.NewDescribeDBClustersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DBClusters {

			result = append(result, Resource{
				Type:      "aws_rds_cluster",
				ID:        *r.DBClusterIdentifier,
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
