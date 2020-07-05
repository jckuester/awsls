// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func ListRoute53ResolverEndpoint(client *Client) ([]Resource, error) {
	req := client.Route53resolverconn.ListResolverEndpointsRequest(&route53resolver.ListResolverEndpointsInput{})

	var result []Resource

	p := route53resolver.NewListResolverEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ResolverEndpoints {

			t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.CreationTime)
			if err != nil {
				return nil, err
			}
			result = append(result, Resource{
				Type:   "aws_route53_resolver_endpoint",
				ID:     *r.Id,
				Region: client.Route53resolverconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
