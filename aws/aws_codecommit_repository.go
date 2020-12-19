// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func ListCodecommitRepository(client *Client) ([]Resource, error) {
	req := client.Codecommitconn.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})

	var result []Resource

	p := codecommit.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Repositories {

			result = append(result, Resource{
				Type:      "aws_codecommit_repository",
				ID:        *r.RepositoryName,
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
