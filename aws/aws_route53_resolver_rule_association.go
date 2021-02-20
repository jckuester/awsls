// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRoute53ResolverRuleAssociation(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Route53resolverconn.ListResolverRuleAssociationsRequest(&route53resolver.ListResolverRuleAssociationsInput{})

	var result []terraform.Resource

	p := route53resolver.NewListResolverRuleAssociationsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ResolverRuleAssociations {

			result = append(result, terraform.Resource{
				Type:      "aws_route53_resolver_rule_association",
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
