// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsReplicationSubnetGroup(client *Client) {
	req := client.databasemigrationserviceconn.DescribeReplicationSubnetGroupsRequest(&databasemigrationservice.DescribeReplicationSubnetGroupsInput{})

	p := databasemigrationservice.NewDescribeReplicationSubnetGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_dms_replication_subnet_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationSubnetGroups {
			fmt.Println(*r.ReplicationSubnetGroupIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_dms_replication_subnet_group: %s", err)
	}

}
