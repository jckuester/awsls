// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

func ListElasticacheReplicationGroup(client *Client) ([]Resource, error) {
	req := client.elasticacheconn.DescribeReplicationGroupsRequest(&elasticache.DescribeReplicationGroupsInput{})

	var result []Resource

	p := elasticache.NewDescribeReplicationGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ReplicationGroups {

			result = append(result, Resource{
				Type: "aws_elasticache_replication_group",
				ID:   *r.ReplicationGroupId,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
