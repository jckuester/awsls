// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

func ListAppsyncGraphqlApi(client *Client) error {
	req := client.appsyncconn.ListGraphqlApisRequest(&appsync.ListGraphqlApisInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.GraphqlApis) > 0 {
		for _, r := range resp.GraphqlApis {
			fmt.Println(*r.ApiId)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	return nil
}
