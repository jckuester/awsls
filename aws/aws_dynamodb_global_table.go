// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ListDynamodbGlobalTable(client *Client) error {
	req := client.dynamodbconn.ListGlobalTablesRequest(&dynamodb.ListGlobalTablesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.GlobalTables) > 0 {
		for _, r := range resp.GlobalTables {
			fmt.Println(*r.GlobalTableName)
		}
	}

	return nil
}
