// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func ListRoute53ResolverRule(client *Client) ([]Resource, error) {
	req := client.Route53resolverconn.ListResolverRulesRequest(&route53resolver.ListResolverRulesInput{})

	var result []Resource

	p := route53resolver.NewListResolverRulesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ResolverRules {
			if *r.OwnerId != client.AccountID {
				continue
			}

			result = append(result, Resource{
				Type:      "aws_route53_resolver_rule",
				ID:        *r.Id,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
