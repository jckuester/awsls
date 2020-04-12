// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamInstanceProfile(client *Client) {
	req := client.iamconn.ListInstanceProfilesRequest(&iam.ListInstanceProfilesInput{})

	p := iam.NewListInstanceProfilesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_iam_instance_profile:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.InstanceProfiles {
			fmt.Println(*r.InstanceProfileName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_iam_instance_profile: %s", err)
	}

}
