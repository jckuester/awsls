// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerCodeRepository(client *Client) ([]Resource, error) {
	req := client.Sagemakerconn.ListCodeRepositoriesRequest(&sagemaker.ListCodeRepositoriesInput{})

	var result []Resource

	p := sagemaker.NewListCodeRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.CodeRepositorySummaryList {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_sagemaker_code_repository",
				ID:        *r.CodeRepositoryName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
