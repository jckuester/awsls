// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftCluster(client *Client) {
	req := client.redshiftconn.DescribeClustersRequest(&redshift.DescribeClustersInput{})

	p := redshift.NewDescribeClustersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_redshift_cluster:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Clusters {
			fmt.Println(*r.ClusterIdentifier)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_redshift_cluster: %s", err)
	}

}
