// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSagemakerCodeRepository(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := sagemaker.NewListCodeRepositoriesPaginator(client.Sagemakerconn, &sagemaker.ListCodeRepositoriesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.CodeRepositorySummaryList {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_sagemaker_code_repository",
				ID:        *r.CodeRepositoryName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
