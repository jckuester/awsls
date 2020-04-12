// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func ListElastictranscoderPipeline(client *Client) error {
	req := client.elastictranscoderconn.ListPipelinesRequest(&elastictranscoder.ListPipelinesInput{})

	p := elastictranscoder.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Pipelines {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
