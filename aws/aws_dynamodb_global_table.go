// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ListDynamodbGlobalTable(client *Client) ([]Resource, error) {
	req := client.Dynamodbconn.ListGlobalTablesRequest(&dynamodb.ListGlobalTablesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.GlobalTables) > 0 {
		for _, r := range resp.GlobalTables {

			result = append(result, Resource{
				Type:   "aws_dynamodb_global_table",
				ID:     *r.GlobalTableName,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
