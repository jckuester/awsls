// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListVpcEndpointService(client *Client) {
	req := client.ec2conn.DescribeVpcEndpointServicesRequest(&ec2.DescribeVpcEndpointServicesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_vpc_endpoint_service: %s", err)
	} else {
		if len(resp.ServiceDetails) > 0 {
			fmt.Println("")
			fmt.Println("aws_vpc_endpoint_service:")
			for _, r := range resp.ServiceDetails {
				fmt.Println(*r.ServiceId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
