// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftCluster(client *Client) ([]Resource, error) {
	req := client.redshiftconn.DescribeClustersRequest(&redshift.DescribeClustersInput{})

	var result []Resource

	p := redshift.NewDescribeClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Clusters {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_redshift_cluster",
				ID:     *r.ClusterIdentifier,
				Region: client.redshiftconn.Config.Region,
				Tags:   tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
