// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListBatchComputeEnvironment(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := batch.NewDescribeComputeEnvironmentsPaginator(client.Batchconn, &batch.DescribeComputeEnvironmentsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.ComputeEnvironments {

			result = append(result, terraform.Resource{
				Type:      "aws_batch_compute_environment",
				ID:        *r.ComputeEnvironmentName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
