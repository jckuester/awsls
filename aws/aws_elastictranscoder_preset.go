// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func ListElastictranscoderPreset(client *Client) ([]Resource, error) {
	req := client.elastictranscoderconn.ListPresetsRequest(&elastictranscoder.ListPresetsInput{})

	var result []Resource

	p := elastictranscoder.NewListPresetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Presets {

			result = append(result, Resource{
				Type:   "aws_elastictranscoder_preset",
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
