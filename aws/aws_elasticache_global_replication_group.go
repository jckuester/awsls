// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElasticacheGlobalReplicationGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := elasticache.NewDescribeGlobalReplicationGroupsPaginator(client.Elasticacheconn, &elasticache.DescribeGlobalReplicationGroupsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.GlobalReplicationGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_elasticache_global_replication_group",
				ID:        *r.GlobalReplicationGroupId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
