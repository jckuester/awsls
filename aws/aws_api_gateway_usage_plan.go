// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func ListApiGatewayUsagePlan(client *Client) error {
	req := client.apigatewayconn.GetUsagePlansRequest(&apigateway.GetUsagePlansInput{})

	p := apigateway.NewGetUsagePlansPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Items {
			fmt.Println(*r.Id)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
