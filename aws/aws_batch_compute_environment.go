// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchComputeEnvironment(client *Client) error {
	req := client.batchconn.DescribeComputeEnvironmentsRequest(&batch.DescribeComputeEnvironmentsInput{})

	p := batch.NewDescribeComputeEnvironmentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ComputeEnvironments {
			fmt.Println(*r.ComputeEnvironmentName)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
