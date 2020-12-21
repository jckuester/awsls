// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

func ListCodepipelineWebhook(client *Client) ([]Resource, error) {
	req := client.Codepipelineconn.ListWebhooksRequest(&codepipeline.ListWebhooksInput{})

	var result []Resource

	p := codepipeline.NewListWebhooksPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Webhooks {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:      "aws_codepipeline_webhook",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
