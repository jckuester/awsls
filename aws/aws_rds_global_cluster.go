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
		page := p.CurrentPage()

		for _, r := range page.GlobalClusters {

			result = append(result, Resource{
				Type:   "aws_rds_global_cluster",
				ID:     *r.GlobalClusterIdentifier,
				Region: client.Rdsconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
