// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

func ListAccessanalyzerAnalyzer(client *Client) ([]Resource, error) {
	req := client.Accessanalyzerconn.ListAnalyzersRequest(&accessanalyzer.ListAnalyzersInput{})

	var result []Resource

	p := accessanalyzer.NewListAnalyzersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Analyzers {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type:    "aws_accessanalyzer_analyzer",
				ID:      *r.Name,
				Profile: client.Profile,
				Region:  client.Region,
				Tags:    tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
