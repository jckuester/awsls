// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsReplicationTask(client *Client) ([]Resource, error) {
	req := client.Databasemigrationserviceconn.DescribeReplicationTasksRequest(&databasemigrationservice.DescribeReplicationTasksInput{})

	var result []Resource

	p := databasemigrationservice.NewDescribeReplicationTasksPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ReplicationTasks {

			result = append(result, Resource{
				Type:      "aws_dms_replication_task",
				ID:        *r.ReplicationTaskIdentifier,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
