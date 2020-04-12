// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func ListRoute53ResolverRuleAssociation(client *Client) {
	req := client.route53resolverconn.ListResolverRuleAssociationsRequest(&route53resolver.ListResolverRuleAssociationsInput{})

	p := route53resolver.NewListResolverRuleAssociationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_route53_resolver_rule_association:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ResolverRuleAssociations {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_route53_resolver_rule_association: %s", err)
	}

}
