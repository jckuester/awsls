// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

func ListAccessanalyzerAnalyzer(client *Client) {
	req := client.accessanalyzerconn.ListAnalyzersRequest(&accessanalyzer.ListAnalyzersInput{})

	p := accessanalyzer.NewListAnalyzersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_accessanalyzer_analyzer:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Analyzers {
			fmt.Println(*r.Name)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_accessanalyzer_analyzer: %s", err)
	}

}
