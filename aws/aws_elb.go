// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElb(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Elasticloadbalancingconn.DescribeLoadBalancersRequest(&elasticloadbalancing.DescribeLoadBalancersInput{})

	var result []terraform.Resource

	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

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

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
