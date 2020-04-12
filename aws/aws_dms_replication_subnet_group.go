// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsReplicationSubnetGroup(client *Client) error {
	req := client.databasemigrationserviceconn.DescribeReplicationSubnetGroupsRequest(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{})

	p := databasemigrationservice.NewDescribeReplicationSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationSubnetGroups {
			fmt.Println(*r.ReplicationSubnetGroupIdentifier)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
