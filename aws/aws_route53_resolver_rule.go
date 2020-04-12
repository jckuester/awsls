// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

func ListRoute53ResolverRule(client *Client) error {
	req := client.route53resolverconn.ListResolverRulesRequest(&route53resolver.ListResolverRulesInput{})

	p := route53resolver.NewListResolverRulesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ResolverRules {
			fmt.Println(*r.Id)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
