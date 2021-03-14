// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListVpcEndpoint(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeVpcEndpointsPaginator(client.Ec2conn, &ec2.DescribeVpcEndpointsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.VpcEndpoints {
			if *r.OwnerId != client.AccountID {
				continue
			}
			t := *r.CreationTimestamp
			result = append(result, terraform.Resource{
				Type:      "aws_vpc_endpoint",
				ID:        *r.VpcEndpointId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
