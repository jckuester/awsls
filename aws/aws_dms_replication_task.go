// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsReplicationTask(client *Client) {
	req := client.databasemigrationserviceconn.DescribeReplicationTasksRequest(&databasemigrationservice.DescribeReplicationTasksInput{})

	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(req)
	fmt.Println("")
	fmt.Println("aws_dms_replication_task:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationTasks {
			fmt.Println(*r.ReplicationTaskIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_dms_replication_task: %s", err)
	}

}
