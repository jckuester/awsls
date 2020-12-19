// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

func ListAppsyncGraphqlApi(client *Client) ([]Resource, error) {
	req := client.Appsyncconn.ListGraphqlApisRequest(&appsync.ListGraphqlApisInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.GraphqlApis) > 0 {

		for _, r := range resp.GraphqlApis {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type:      "aws_appsync_graphql_api",
				ID:        *r.ApiId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	return result, nil
}
