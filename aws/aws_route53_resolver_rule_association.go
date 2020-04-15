// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func ListRoute53ResolverRuleAssociation(client *Client) ([]Resource, error) {
	req := client.route53resolverconn.ListResolverRuleAssociationsRequest(&route53resolver.ListResolverRuleAssociationsInput{})

	var result []Resource

	p := route53resolver.NewListResolverRuleAssociationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ResolverRuleAssociations {

			result = append(result, Resource{
				Type: "aws_route53_resolver_rule_association",
				ID:   *r.Id,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
