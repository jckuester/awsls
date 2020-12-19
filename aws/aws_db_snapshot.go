// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbSnapshot(client *Client) ([]Resource, error) {
	req := client.Rdsconn.DescribeDBSnapshotsRequest(&rds.DescribeDBSnapshotsInput{})

	var result []Resource

	p := rds.NewDescribeDBSnapshotsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DBSnapshots {

			t := *r.InstanceCreateTime
			result = append(result, Resource{
				Type:      "aws_db_snapshot",
				ID:        *r.DBSnapshotIdentifier,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
