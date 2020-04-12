// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
)

func ListElastictranscoderPreset(client *Client) {
	req := client.elastictranscoderconn.ListPresetsRequest(&elastictranscoder.ListPresetsInput{})

	p := elastictranscoder.NewListPresetsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_elastictranscoder_preset:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Presets {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_elastictranscoder_preset: %s", err)
	}

}
