// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListOpsworksUserProfile(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Opsworksconn.DescribeUserProfiles(ctx, &opsworks.DescribeUserProfilesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.UserProfiles) > 0 {

		for _, r := range resp.UserProfiles {

			result = append(result, terraform.Resource{
				Type:      "aws_opsworks_user_profile",
				ID:        *r.IamUserArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
