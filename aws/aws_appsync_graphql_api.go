// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAppsyncGraphqlApi(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Appsyncconn.ListGraphqlApis(ctx, &appsync.ListGraphqlApisInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.GraphqlApis) > 0 {

		for _, r := range resp.GraphqlApis {

			result = append(result, terraform.Resource{
				Type:      "aws_appsync_graphql_api",
				ID:        *r.ApiId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
