// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchJobQueue(client *Client) error {
	req := client.batchconn.DescribeJobQueuesRequest(&batch.DescribeJobQueuesInput{})

	p := batch.NewDescribeJobQueuesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.JobQueues {
			fmt.Println(*r.JobQueueArn)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
