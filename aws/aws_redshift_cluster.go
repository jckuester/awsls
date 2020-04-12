// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftCluster(client *Client) error {
	req := client.redshiftconn.DescribeClustersRequest(&redshift.DescribeClustersInput{})

	p := redshift.NewDescribeClustersPaginator(req)
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
		return err
	}

	return nil
}
