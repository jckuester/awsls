// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEbsSnapshot(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := ec2.NewDescribeSnapshotsPaginator(client.Ec2conn, &ec2.DescribeSnapshotsInput{
		OwnerIds: []string{"self"},
	})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Snapshots {
			if *r.OwnerId != client.AccountID {
				continue
			}
			t := *r.StartTime
			result = append(result, terraform.Resource{
				Type:      "aws_ebs_snapshot",
				ID:        *r.SnapshotId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
