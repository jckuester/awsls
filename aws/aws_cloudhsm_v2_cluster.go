// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

func ListCloudhsmV2Cluster(client *Client) ([]Resource, error) {
	req := client.Cloudhsmv2conn.DescribeClustersRequest(&cloudhsmv2.DescribeClustersInput{})

	var result []Resource

	p := cloudhsmv2.NewDescribeClustersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Clusters {

			result = append(result, Resource{
				Type:      "aws_cloudhsm_v2_cluster",
				ID:        *r.ClusterId,
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
