// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchJobQueue(client *Client) ([]Resource, error) {
	req := client.Batchconn.DescribeJobQueuesRequest(&batch.DescribeJobQueuesInput{})

	var result []Resource

	p := batch.NewDescribeJobQueuesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.JobQueues {

			result = append(result, Resource{
				Type:    "aws_batch_job_queue",
				ID:      *r.JobQueueArn,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
