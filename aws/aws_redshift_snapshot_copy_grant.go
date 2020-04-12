// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftSnapshotCopyGrant(client *Client) error {
	req := client.redshiftconn.DescribeSnapshotCopyGrantsRequest(&redshift.DescribeSnapshotCopyGrantsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.SnapshotCopyGrants) > 0 {
		for _, r := range resp.SnapshotCopyGrants {
			fmt.Println(*r.SnapshotCopyGrantName)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
