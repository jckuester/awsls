// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksUserProfile(client *Client) ([]Resource, error) {
	req := client.Opsworksconn.DescribeUserProfilesRequest(&opsworks.DescribeUserProfilesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.UserProfiles) > 0 {
		for _, r := range resp.UserProfiles {

			result = append(result, Resource{
				Type:   "aws_opsworks_user_profile",
				ID:     *r.IamUserArn,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
