// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func ListEcsCluster(client *Client) {
	req := client.ecsconn.DescribeClustersRequest(&ecs.DescribeClustersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_ecs_cluster: %s", err)
	} else {
		if len(resp.Clusters) > 0 {
			fmt.Println("")
			fmt.Println("aws_ecs_cluster:")
			for _, r := range resp.Clusters {
				fmt.Println(*r.ClusterArn)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
