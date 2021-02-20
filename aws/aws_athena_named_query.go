// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAthenaNamedQuery(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Athenaconn.ListNamedQueriesRequest(&athena.ListNamedQueriesInput{})

	var result []terraform.Resource

	p := athena.NewListNamedQueriesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.NamedQueryIds {

			result = append(result, terraform.Resource{
				Type:      "aws_athena_named_query",
				ID:        r,
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
