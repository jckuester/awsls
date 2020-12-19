// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueTrigger(client *Client) ([]Resource, error) {
	req := client.Glueconn.GetTriggersRequest(&glue.GetTriggersInput{})

	var result []Resource

	p := glue.NewGetTriggersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Triggers {

			result = append(result, Resource{
				Type:      "aws_glue_trigger",
				ID:        *r.Name,
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
