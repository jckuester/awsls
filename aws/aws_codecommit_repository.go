// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func ListCodecommitRepository(client *Client) ([]Resource, error) {
	req := client.codecommitconn.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})

	var result []Resource

	p := codecommit.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Repositories {

			result = append(result, Resource{
				Type:   "aws_codecommit_repository",
				ID:     *r.RepositoryName,
				Region: client.codecommitconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
