// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListVpcEndpointService(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Ec2conn.DescribeVpcEndpointServices(ctx, &ec2.DescribeVpcEndpointServicesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ServiceDetails) > 0 {

		for _, r := range resp.ServiceDetails {

			result = append(result, terraform.Resource{
				Type:      "aws_vpc_endpoint_service",
				ID:        *r.ServiceId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
