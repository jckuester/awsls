// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueTrigger(client *Client) ([]Resource, error) {
	req := client.glueconn.GetTriggersRequest(&glue.GetTriggersInput{})

	var result []Resource

	p := glue.NewGetTriggersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Triggers {

			result = append(result, Resource{
				Type:   "aws_glue_trigger",
				ID:     *r.Name,
				Region: client.glueconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
