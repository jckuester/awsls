// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func ListMskCluster(client *Client) {
	req := client.kafkaconn.ListClustersRequest(&kafka.ListClustersInput{})

	p := kafka.NewListClustersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_msk_cluster:")
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
		log.Printf("aws_msk_cluster: %s", err)
	}

}
