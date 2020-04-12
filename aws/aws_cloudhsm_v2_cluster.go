// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

func ListCloudhsmV2Cluster(client *Client) {
	req := client.cloudhsmv2conn.DescribeClustersRequest(&cloudhsmv2.DescribeClustersInput{})

	p := cloudhsmv2.NewDescribeClustersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_cloudhsm_v2_cluster:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Clusters {
			fmt.Println(*r.ClusterId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_cloudhsm_v2_cluster: %s", err)
	}

}
