// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSecurityhubActionTarget(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := securityhub.NewDescribeActionTargetsPaginator(client.Securityhubconn, &securityhub.DescribeActionTargetsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.ActionTargets {

			result = append(result, terraform.Resource{
				Type:      "aws_securityhub_action_target",
				ID:        *r.ActionTargetArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
