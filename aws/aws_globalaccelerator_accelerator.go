// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGlobalacceleratorAccelerator(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := globalaccelerator.NewListAcceleratorsPaginator(client.Globalacceleratorconn, &globalaccelerator.ListAcceleratorsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Accelerators {

			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_globalaccelerator_accelerator",
				ID:        *r.AcceleratorArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
