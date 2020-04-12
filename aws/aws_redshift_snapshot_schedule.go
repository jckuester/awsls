// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftSnapshotSchedule(client *Client) {
	req := client.redshiftconn.DescribeSnapshotSchedulesRequest(&redshift.DescribeSnapshotSchedulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_redshift_snapshot_schedule: %s", err)
	} else {
		if len(resp.SnapshotSchedules) > 0 {
			fmt.Println("")
			fmt.Println("aws_redshift_snapshot_schedule:")
			for _, r := range resp.SnapshotSchedules {
				fmt.Println(*r.ScheduleIdentifier)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
