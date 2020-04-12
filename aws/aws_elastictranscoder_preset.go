// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func ListElastictranscoderPreset(client *Client) error {
	req := client.elastictranscoderconn.ListPresetsRequest(&elastictranscoder.ListPresetsInput{})

	p := elastictranscoder.NewListPresetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Presets {
			fmt.Println(*r.Id)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
