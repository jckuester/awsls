// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ListDynamodbGlobalTable(client *Client) {
	req := client.dynamodbconn.ListGlobalTablesRequest(&dynamodb.ListGlobalTablesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_dynamodb_global_table: %s", err)
	} else {
		if len(resp.GlobalTables) > 0 {
			fmt.Println("")
			fmt.Println("aws_dynamodb_global_table:")
			for _, r := range resp.GlobalTables {
				fmt.Println(*r.GlobalTableName)

			}
		}
	}

}
