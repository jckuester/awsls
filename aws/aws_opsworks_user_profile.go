// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksUserProfile(client *Client) {
	req := client.opsworksconn.DescribeUserProfilesRequest(&opsworks.DescribeUserProfilesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_opsworks_user_profile: %s", err)
	} else {
		if len(resp.UserProfiles) > 0 {
			fmt.Println("")
			fmt.Println("aws_opsworks_user_profile:")
			for _, r := range resp.UserProfiles {
				fmt.Println(*r.IamUserArn)

			}
		}
	}

}
