// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

func ListAppsyncGraphqlApi(client *Client) {
	req := client.appsyncconn.ListGraphqlApisRequest(&appsync.ListGraphqlApisInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_appsync_graphql_api: %s", err)
	} else {
		if len(resp.GraphqlApis) > 0 {
			fmt.Println("")
			fmt.Println("aws_appsync_graphql_api:")
			for _, r := range resp.GraphqlApis {
				fmt.Println(*r.ApiId)
				for k, v := range r.Tags {
					fmt.Printf("\t%s: %s\n", k, v)
				}
			}
		}
	}

}
