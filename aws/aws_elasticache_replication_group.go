// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

func ListElasticacheReplicationGroup(client *Client) error {
	req := client.elasticacheconn.DescribeReplicationGroupsRequest(&elasticache.DescribeReplicationGroupsInput{})

	p := elasticache.NewDescribeReplicationGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationGroups {
			fmt.Println(*r.ReplicationGroupId)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
