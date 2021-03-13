// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudhsmV2Cluster(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := cloudhsmv2.NewDescribeClustersPaginator(client.Cloudhsmv2conn, &cloudhsmv2.DescribeClustersInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Clusters {

			result = append(result, terraform.Resource{
				Type:      "aws_cloudhsm_v2_cluster",
				ID:        *r.ClusterId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
