// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
)

func ListCloudhsmV2Cluster(client *Client) error {
	req := client.cloudhsmv2conn.DescribeClustersRequest(&cloudhsmv2.DescribeClustersInput{})

	p := cloudhsmv2.NewDescribeClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Clusters {
			fmt.Println(*r.ClusterId)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
