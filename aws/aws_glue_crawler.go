// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueCrawler(client *Client) ([]Resource, error) {
	req := client.Glueconn.GetCrawlersRequest(&glue.GetCrawlersInput{})

	var result []Resource

	p := glue.NewGetCrawlersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Crawlers {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:   "aws_glue_crawler",
				ID:     *r.Name,
				Region: client.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
