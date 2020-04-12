// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func ListEcrRepository(client *Client) error {
	req := client.ecrconn.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{})

	p := ecr.NewDescribeRepositoriesPaginator(req)
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
