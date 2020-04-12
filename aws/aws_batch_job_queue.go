// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchJobQueue(client *Client) {
	req := client.batchconn.DescribeJobQueuesRequest(&batch.DescribeJobQueuesInput{})

	p := batch.NewDescribeJobQueuesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_batch_job_queue:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.JobQueues {
			fmt.Println(*r.JobQueueArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_batch_job_queue: %s", err)
	}

}
