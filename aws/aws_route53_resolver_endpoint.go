// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRoute53ResolverEndpoint(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := route53resolver.NewListResolverEndpointsPaginator(client.Route53resolverconn, &route53resolver.ListResolverEndpointsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.ResolverEndpoints {

			t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.CreationTime)
			if err != nil {
				return nil, err
			}
			result = append(result, terraform.Resource{
				Type:      "aws_route53_resolver_endpoint",
				ID:        *r.Id,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
