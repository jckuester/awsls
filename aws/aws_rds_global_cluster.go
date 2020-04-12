// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListRdsGlobalCluster(client *Client) {
	req := client.rdsconn.DescribeGlobalClustersRequest(&rds.DescribeGlobalClustersInput{})

	p := rds.NewDescribeGlobalClustersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_rds_global_cluster:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.GlobalClusters {
			fmt.Println(*r.GlobalClusterIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_rds_global_cluster: %s", err)
	}

}
