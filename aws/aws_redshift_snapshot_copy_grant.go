// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftSnapshotCopyGrant(client *Client) ([]Resource, error) {
	req := client.Redshiftconn.DescribeSnapshotCopyGrantsRequest(&redshift.DescribeSnapshotCopyGrantsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SnapshotCopyGrants) > 0 {
		for _, r := range resp.SnapshotCopyGrants {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:      "aws_redshift_snapshot_copy_grant",
				ID:        *r.SnapshotCopyGrantName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	return result, nil
}
