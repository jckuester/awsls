// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDynamodbTable(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Dynamodbconn.ListTablesRequest(&dynamodb.ListTablesInput{})

	var result []terraform.Resource

	p := dynamodb.NewListTablesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.TableNames {

			result = append(result, terraform.Resource{
				Type:      "aws_dynamodb_table",
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
