// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEbsSnapshot(client *Client) {
	req := client.ec2conn.DescribeSnapshotsRequest(&ec2.DescribeSnapshotsInput{})

	p := ec2.NewDescribeSnapshotsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ebs_snapshot:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Snapshots {
			fmt.Println(*r.SnapshotId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ebs_snapshot: %s", err)
	}

}
