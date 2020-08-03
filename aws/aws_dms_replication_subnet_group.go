// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsReplicationSubnetGroup(client *Client) ([]Resource, error) {
	req := client.Databasemigrationserviceconn.DescribeReplicationSubnetGroupsRequest(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{})

	var result []Resource

	p := databasemigrationservice.NewDescribeReplicationSubnetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationSubnetGroups {

			result = append(result, Resource{
				Type:      "aws_dms_replication_subnet_group",
				ID:        *r.ReplicationSubnetGroupIdentifier,
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
