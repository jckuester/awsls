// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func ListMskCluster(client *Client) error {
	req := client.kafkaconn.ListClustersRequest(&kafka.ListClustersInput{})

	p := kafka.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ClusterInfoList {
			fmt.Println(*r.ClusterArn)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
