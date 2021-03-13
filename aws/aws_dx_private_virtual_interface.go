// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDxPrivateVirtualInterface(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Directconnectconn.DescribeVirtualInterfaces(ctx, &directconnect.DescribeVirtualInterfacesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.VirtualInterfaces) > 0 {

		for _, r := range resp.VirtualInterfaces {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, terraform.Resource{
				Type:      "aws_dx_private_virtual_interface",
				ID:        *r.VirtualInterfaceId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	return result, nil
}
