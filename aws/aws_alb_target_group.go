// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListAlbTargetGroup(client *Client) {
	req := client.elasticloadbalancingv2conn.DescribeTargetGroupsRequest(&elasticloadbalancingv2.DescribeTargetGroupsInput{})

	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_alb_target_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TargetGroups {
			fmt.Println(*r.TargetGroupArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_alb_target_group: %s", err)
	}

}
