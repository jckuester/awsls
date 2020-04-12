// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftSnapshotSchedule(client *Client) error {
	req := client.redshiftconn.DescribeSnapshotSchedulesRequest(&redshift.DescribeSnapshotSchedulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.SnapshotSchedules) > 0 {
		for _, r := range resp.SnapshotSchedules {
			fmt.Println(*r.ScheduleIdentifier)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
