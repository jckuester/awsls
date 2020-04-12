// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

func ListElb(client *Client) {
	req := client.elasticloadbalancingconn.DescribeLoadBalancersRequest(&elasticloadbalancing.DescribeLoadBalancersInput{})

	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_elb:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LoadBalancerDescriptions {
			fmt.Println(*r.LoadBalancerName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_elb: %s", err)
	}

}
