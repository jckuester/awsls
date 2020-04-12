// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func ListElastictranscoderPipeline(client *Client) {
	req := client.elastictranscoderconn.ListPipelinesRequest(&elastictranscoder.ListPipelinesInput{})

	p := elastictranscoder.NewListPipelinesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_elastictranscoder_pipeline:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Pipelines {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_elastictranscoder_pipeline: %s", err)
	}

}
