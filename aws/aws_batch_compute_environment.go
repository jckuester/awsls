// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchComputeEnvironment(client *Client) {
	req := client.batchconn.DescribeComputeEnvironmentsRequest(&batch.DescribeComputeEnvironmentsInput{})

	p := batch.NewDescribeComputeEnvironmentsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_batch_compute_environment:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ComputeEnvironments {
			fmt.Println(*r.ComputeEnvironmentName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_batch_compute_environment: %s", err)
	}

}
