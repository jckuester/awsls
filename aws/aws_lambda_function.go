// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func ListLambdaFunction(client *Client) {
	req := client.lambdaconn.ListFunctionsRequest(&lambda.ListFunctionsInput{})

	p := lambda.NewListFunctionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_lambda_function:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Functions {
			fmt.Println(*r.FunctionName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_lambda_function: %s", err)
	}

}
