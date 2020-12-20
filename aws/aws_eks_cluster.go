// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func ListEksCluster(client *Client) ([]Resource, error) {
	req := client.Eksconn.ListClustersRequest(&eks.ListClustersInput{})

	var result []Resource

	p := eks.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Clusters {

			result = append(result, Resource{
				Type:      "aws_eks_cluster",
				ID:        r,
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
