// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLambdaCodeSigningConfig(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := lambda.NewListCodeSigningConfigsPaginator(client.Lambdaconn, &lambda.ListCodeSigningConfigsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.CodeSigningConfigs {

			result = append(result, terraform.Resource{
				Type:      "aws_lambda_code_signing_config",
				ID:        *r.CodeSigningConfigArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
