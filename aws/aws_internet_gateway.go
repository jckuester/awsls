// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListInternetGateway(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeInternetGatewaysPaginator(client.Ec2conn, &ec2.DescribeInternetGatewaysInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.InternetGateways {
			if *r.OwnerId != client.AccountID {
				continue
			}

			result = append(result, terraform.Resource{
				Type:      "aws_internet_gateway",
				ID:        *r.InternetGatewayId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
