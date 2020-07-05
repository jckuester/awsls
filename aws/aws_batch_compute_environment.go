// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchComputeEnvironment(client *Client) ([]Resource, error) {
	req := client.Batchconn.DescribeComputeEnvironmentsRequest(&batch.DescribeComputeEnvironmentsInput{})

	var result []Resource

	p := batch.NewDescribeComputeEnvironmentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ComputeEnvironments {

			result = append(result, Resource{
				Type:   "aws_batch_compute_environment",
				ID:     *r.ComputeEnvironmentName,
				Region: client.Batchconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
