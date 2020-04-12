// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/batch"
)

func ListBatchJobDefinition(client *Client) error {
	req := client.batchconn.DescribeJobDefinitionsRequest(&batch.DescribeJobDefinitionsInput{})

	p := batch.NewDescribeJobDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.JobDefinitions {
			fmt.Println(*r.JobDefinitionArn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
