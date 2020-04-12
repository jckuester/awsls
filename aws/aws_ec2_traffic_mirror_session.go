// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TrafficMirrorSession(client *Client) {
	req := client.ec2conn.DescribeTrafficMirrorSessionsRequest(&ec2.DescribeTrafficMirrorSessionsInput{})

	p := ec2.NewDescribeTrafficMirrorSessionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ec2_traffic_mirror_session:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TrafficMirrorSessions {
			fmt.Println(*r.TrafficMirrorSessionId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_ec2_traffic_mirror_session: %s", err)
	}

}
