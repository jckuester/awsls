// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func ListEcsCluster(client *Client) ([]Resource, error) {
	req := client.Ecsconn.DescribeClustersRequest(&ecs.DescribeClustersInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Clusters) > 0 {
		for _, r := range resp.Clusters {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_ecs_cluster",
				ID:     *r.ClusterArn,
				Region: client.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
