// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftSnapshotCopyGrant(client *Client) {
	req := client.redshiftconn.DescribeSnapshotCopyGrantsRequest(&redshift.DescribeSnapshotCopyGrantsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_redshift_snapshot_copy_grant: %s", err)
	} else {
		if len(resp.SnapshotCopyGrants) > 0 {
			fmt.Println("")
			fmt.Println("aws_redshift_snapshot_copy_grant:")
			for _, r := range resp.SnapshotCopyGrants {
				fmt.Println(*r.SnapshotCopyGrantName)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
