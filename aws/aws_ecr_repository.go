// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func ListEcrRepository(client *Client) ([]Resource, error) {
	req := client.ecrconn.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{})

	var result []Resource

	p := ecr.NewDescribeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Repositories {

			result = append(result, Resource{
				Type: "aws_ecr_repository",
				ID:   *r.RepositoryName,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
