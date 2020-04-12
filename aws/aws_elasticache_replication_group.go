// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

func ListElasticacheReplicationGroup(client *Client) {
	req := client.elasticacheconn.DescribeReplicationGroupsRequest(&elasticache.DescribeReplicationGroupsInput{})

	p := elasticache.NewDescribeReplicationGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_elasticache_replication_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationGroups {
			fmt.Println(*r.ReplicationGroupId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_elasticache_replication_group: %s", err)
	}

}
