// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListRdsGlobalCluster(client *Client) error {
	req := client.rdsconn.DescribeGlobalClustersRequest(&rds.DescribeGlobalClustersInput{})

	p := rds.NewDescribeGlobalClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.GlobalClusters {
			fmt.Println(*r.GlobalClusterIdentifier)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
