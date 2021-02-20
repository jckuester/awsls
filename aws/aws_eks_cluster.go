// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEksCluster(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Eksconn.ListClustersRequest(&eks.ListClustersInput{})

	var result []terraform.Resource

	p := eks.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Clusters {

			result = append(result, terraform.Resource{
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
