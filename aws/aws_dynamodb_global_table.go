// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDynamodbGlobalTable(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Dynamodbconn.ListGlobalTables(ctx, &dynamodb.ListGlobalTablesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.GlobalTables) > 0 {

		for _, r := range resp.GlobalTables {

			result = append(result, terraform.Resource{
				Type:      "aws_dynamodb_global_table",
				ID:        *r.GlobalTableName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
