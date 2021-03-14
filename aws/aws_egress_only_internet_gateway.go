// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEgressOnlyInternetGateway(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeEgressOnlyInternetGatewaysPaginator(client.Ec2conn, &ec2.DescribeEgressOnlyInternetGatewaysInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.EgressOnlyInternetGateways {

			result = append(result, terraform.Resource{
				Type:      "aws_egress_only_internet_gateway",
				ID:        *r.EgressOnlyInternetGatewayId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
