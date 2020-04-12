// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func ListEcsCluster(client *Client) error {
	req := client.ecsconn.DescribeClustersRequest(&ecs.DescribeClustersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Clusters) > 0 {
		for _, r := range resp.Clusters {
			fmt.Println(*r.ClusterArn)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
