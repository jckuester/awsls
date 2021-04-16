// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListStoragegatewayTapePool(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := storagegateway.NewListTapePoolsPaginator(client.Storagegatewayconn, &storagegateway.ListTapePoolsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.PoolInfos {

			result = append(result, terraform.Resource{
				Type:      "aws_storagegateway_tape_pool",
				ID:        *r.PoolARN,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
