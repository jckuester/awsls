// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TrafficMirrorTarget(client *Client) {
	req := client.ec2conn.DescribeTrafficMirrorTargetsRequest(&ec2.DescribeTrafficMirrorTargetsInput{})

	p := ec2.NewDescribeTrafficMirrorTargetsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ec2_traffic_mirror_target:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TrafficMirrorTargets {
			fmt.Println(*r.TrafficMirrorTargetId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ec2_traffic_mirror_target: %s", err)
	}

}
