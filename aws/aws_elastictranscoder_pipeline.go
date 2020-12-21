// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func ListElastictranscoderPipeline(client *Client) ([]Resource, error) {
	req := client.Elastictranscoderconn.ListPipelinesRequest(&elastictranscoder.ListPipelinesInput{})

	var result []Resource

	p := elastictranscoder.NewListPipelinesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Pipelines {

			result = append(result, Resource{
				Type:      "aws_elastictranscoder_pipeline",
				ID:        *r.Id,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
