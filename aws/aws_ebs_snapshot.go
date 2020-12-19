// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEbsSnapshot(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{})

	var result []Resource

	p := ec2.NewDescribeSnapshotsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Snapshots {
			if *r.OwnerId != client.AccountID {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.StartTime
			result = append(result, Resource{
				Type:      "aws_ebs_snapshot",
				ID:        *r.SnapshotId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
