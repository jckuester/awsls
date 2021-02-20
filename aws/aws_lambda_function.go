// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLambdaFunction(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Lambdaconn.ListFunctionsRequest(&lambda.ListFunctionsInput{})

	var result []terraform.Resource

	p := lambda.NewListFunctionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Functions {

			result = append(result, terraform.Resource{
				Type:      "aws_lambda_function",
				ID:        *r.FunctionName,
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
