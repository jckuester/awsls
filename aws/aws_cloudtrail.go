// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudtrail(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Cloudtrailconn.DescribeTrails(ctx, &cloudtrail.DescribeTrailsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.TrailList) > 0 {

		for _, r := range resp.TrailList {

			result = append(result, terraform.Resource{
				Type:      "aws_cloudtrail",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
