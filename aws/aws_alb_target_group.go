// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListAlbTargetGroup(client *Client) error {
	req := client.elasticloadbalancingv2conn.DescribeTargetGroupsRequest(&elasticloadbalancingv2.DescribeTargetGroupsInput{})

	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TargetGroups {
			fmt.Println(*r.TargetGroupArn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
