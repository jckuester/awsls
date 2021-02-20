// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEcsCluster(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ecsconn.ListClustersRequest(&ecs.ListClustersInput{})

	var result []terraform.Resource

	p := ecs.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ClusterArns {

			result = append(result, terraform.Resource{
				Type:      "aws_ecs_cluster",
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
