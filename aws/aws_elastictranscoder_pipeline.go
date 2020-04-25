// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func ListElastictranscoderPipeline(client *Client) ([]Resource, error) {
	req := client.elastictranscoderconn.ListPipelinesRequest(&elastictranscoder.ListPipelinesInput{})

	var result []Resource

	p := elastictranscoder.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Pipelines {

			result = append(result, Resource{
				Type:   "aws_elastictranscoder_pipeline",
				ID:     *r.Id,
				Region: client.elastictranscoderconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
