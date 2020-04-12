// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func ListRoute53ResolverEndpoint(client *Client) {
	req := client.route53resolverconn.ListResolverEndpointsRequest(&route53resolver.ListResolverEndpointsInput{})

	p := route53resolver.NewListResolverEndpointsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_route53_resolver_endpoint:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ResolverEndpoints {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_route53_resolver_endpoint: %s", err)
	}

}
