// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func ListRoute53ResolverEndpoint(client *Client) error {
	req := client.route53resolverconn.ListResolverEndpointsRequest(&route53resolver.ListResolverEndpointsInput{})

	p := route53resolver.NewListResolverEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ResolverEndpoints {
			fmt.Println(*r.Id)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
