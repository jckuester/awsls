// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListNetworkfirewallFirewallPolicy(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := networkfirewall.NewListFirewallPoliciesPaginator(client.Networkfirewallconn, &networkfirewall.ListFirewallPoliciesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.FirewallPolicies {

			result = append(result, terraform.Resource{
				Type:      "aws_networkfirewall_firewall_policy",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
