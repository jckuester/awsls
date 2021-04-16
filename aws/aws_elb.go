// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElb(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(client.Elasticloadbalancingconn, &elasticloadbalancing.DescribeLoadBalancersInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.LoadBalancerDescriptions {

			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_elb",
				ID:        *r.LoadBalancerName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
