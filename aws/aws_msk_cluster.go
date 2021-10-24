// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListMskCluster(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := kafka.NewListClustersPaginator(client.Kafkaconn, &kafka.ListClustersInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.ClusterInfoList {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}
			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_msk_cluster",
				ID:        *r.ClusterArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
