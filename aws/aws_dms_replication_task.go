// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsReplicationTask(client *Client) error {
	req := client.databasemigrationserviceconn.DescribeReplicationTasksRequest(&databasemigrationservice.DescribeReplicationTasksInput{})

	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationTasks {
			fmt.Println(*r.ReplicationTaskIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
