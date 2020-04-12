// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueCrawler(client *Client) error {
	req := client.glueconn.GetCrawlersRequest(&glue.GetCrawlersInput{})

	p := glue.NewGetCrawlersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Crawlers {
			fmt.Println(*r.Name)

			fmt.Printf("CreatedAt: %s\n", *r.CreationTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
