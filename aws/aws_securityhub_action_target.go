// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSecurityhubActionTarget(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Securityhubconn.DescribeActionTargetsRequest(&securityhub.DescribeActionTargetsInput{})

	var result []terraform.Resource

	p := securityhub.NewDescribeActionTargetsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

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

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
