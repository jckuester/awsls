// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueCrawler(client *Client) {
	req := client.glueconn.GetCrawlersRequest(&glue.GetCrawlersInput{})

	p := glue.NewGetCrawlersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_glue_crawler:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Crawlers {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_glue_crawler: %s", err)
	}

}
