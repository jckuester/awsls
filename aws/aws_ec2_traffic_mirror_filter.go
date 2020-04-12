// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TrafficMirrorFilter(client *Client) {
	req := client.ec2conn.DescribeTrafficMirrorFiltersRequest(&ec2.DescribeTrafficMirrorFiltersInput{})

	p := ec2.NewDescribeTrafficMirrorFiltersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ec2_traffic_mirror_filter:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TrafficMirrorFilters {
			fmt.Println(*r.TrafficMirrorFilterId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ec2_traffic_mirror_filter: %s", err)
	}

}
