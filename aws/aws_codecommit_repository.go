// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func ListCodecommitRepository(client *Client) error {
	req := client.codecommitconn.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})

	p := codecommit.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Repositories {
			fmt.Println(*r.RepositoryName)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
