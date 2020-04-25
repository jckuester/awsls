// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsReplicationTask(client *Client) ([]Resource, error) {
	req := client.databasemigrationserviceconn.DescribeReplicationTasksRequest(&databasemigrationservice.DescribeReplicationTasksInput{})

	var result []Resource

	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationTasks {

			result = append(result, Resource{
				Type:   "aws_dms_replication_task",
				ID:     *r.ReplicationTaskIdentifier,
				Region: client.databasemigrationserviceconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
