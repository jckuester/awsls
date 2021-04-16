// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEc2CarrierGateway(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeCarrierGatewaysPaginator(client.Ec2conn, &ec2.DescribeCarrierGatewaysInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.CarrierGateways {
			if *r.OwnerId != client.AccountID {
				continue
			}

			result = append(result, terraform.Resource{
				Type:      "aws_ec2_carrier_gateway",
				ID:        *r.CarrierGatewayId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
