// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEfsAccessPoint(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := efs.NewDescribeAccessPointsPaginator(client.Efsconn, &efs.DescribeAccessPointsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.AccessPoints {
			if *r.OwnerId != client.AccountID {
				continue
			}

			result = append(result, terraform.Resource{
				Type:      "aws_efs_access_point",
				ID:        *r.AccessPointId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
