// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func ListLambdaFunction(client *Client) error {
	req := client.lambdaconn.ListFunctionsRequest(&lambda.ListFunctionsInput{})

	p := lambda.NewListFunctionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Functions {
			fmt.Println(*r.FunctionName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
