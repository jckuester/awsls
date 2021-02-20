// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudhsmV2Cluster(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Cloudhsmv2conn.DescribeClustersRequest(&cloudhsmv2.DescribeClustersInput{})

	var result []terraform.Resource

	p := cloudhsmv2.NewDescribeClustersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

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

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
