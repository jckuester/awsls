// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudwatchLogResourcePolicy(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Cloudwatchlogsconn.DescribeResourcePoliciesRequest(&cloudwatchlogs.DescribeResourcePoliciesInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ResourcePolicies) > 0 {

		for _, r := range resp.ResourcePolicies {

			result = append(result, terraform.Resource{
				Type:      "aws_cloudwatch_log_resource_policy",
				ID:        *r.PolicyName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
