// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListBatchJobQueue(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Batchconn.DescribeJobQueuesRequest(&batch.DescribeJobQueuesInput{})

	var result []terraform.Resource

	p := batch.NewDescribeJobQueuesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.JobQueues {

			result = append(result, terraform.Resource{
				Type:      "aws_batch_job_queue",
				ID:        *r.JobQueueArn,
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
