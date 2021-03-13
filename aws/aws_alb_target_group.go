// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAlbTargetGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(client.Elasticloadbalancingv2conn, &elasticloadbalancingv2.DescribeTargetGroupsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.TargetGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_alb_target_group",
				ID:        *r.TargetGroupArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
