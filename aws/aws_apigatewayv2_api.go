// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

func ListApigatewayv2Api(client *Client) error {
	req := client.apigatewayv2conn.GetApisRequest(&apigatewayv2.GetApisInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Items) > 0 {
		for _, r := range resp.Items {
			fmt.Println(*r.ApiId)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	return nil
}
