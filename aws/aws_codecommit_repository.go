// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodecommitRepository(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codecommitconn.ListRepositoriesRequest(&codecommit.ListRepositoriesInput{})

	var result []terraform.Resource

	p := codecommit.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Repositories {

			result = append(result, terraform.Resource{
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
