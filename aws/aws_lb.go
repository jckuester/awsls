// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLb(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client.Elasticloadbalancingv2conn, &elasticloadbalancingv2.DescribeLoadBalancersInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.LoadBalancers {

			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_lb",
				ID:        *r.LoadBalancerArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
