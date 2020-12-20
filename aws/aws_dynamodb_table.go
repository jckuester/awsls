// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ListDynamodbTable(client *Client) ([]Resource, error) {
	req := client.Dynamodbconn.ListTablesRequest(&dynamodb.ListTablesInput{})

	var result []Resource

	p := dynamodb.NewListTablesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.TableNames {

			result = append(result, Resource{
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
