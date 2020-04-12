// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

func ListCodepipelineWebhook(client *Client) {
	req := client.codepipelineconn.ListWebhooksRequest(&codepipeline.ListWebhooksInput{})

	p := codepipeline.NewListWebhooksPaginator(req)
	fmt.Println("")
	fmt.Println("aws_codepipeline_webhook:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Webhooks {
			fmt.Println(*r.Arn)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_codepipeline_webhook: %s", err)
	}

}
