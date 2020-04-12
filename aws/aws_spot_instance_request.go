// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListSpotInstanceRequest(client *Client) {
	req := client.ec2conn.DescribeSpotInstanceRequestsRequest(&ec2.DescribeSpotInstanceRequestsInput{})

	p := ec2.NewDescribeSpotInstanceRequestsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_spot_instance_request:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SpotInstanceRequests {
			fmt.Println(*r.SpotInstanceRequestId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_spot_instance_request: %s", err)
	}

}
