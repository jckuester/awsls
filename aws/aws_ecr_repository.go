// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEcrRepository(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ecrconn.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{})

	var result []terraform.Resource

	p := ecr.NewDescribeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Repositories {

			result = append(result, terraform.Resource{
				Type:      "aws_ecr_repository",
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
