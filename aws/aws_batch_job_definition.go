// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchJobDefinition(client *Client) {
	req := client.batchconn.DescribeJobDefinitionsRequest(&batch.DescribeJobDefinitionsInput{})

	p := batch.NewDescribeJobDefinitionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_batch_job_definition:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.JobDefinitions {
			fmt.Println(*r.JobDefinitionArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_batch_job_definition: %s", err)
	}

}
