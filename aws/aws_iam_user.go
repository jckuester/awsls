// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamUser(client *Client) {
	req := client.iamconn.ListUsersRequest(&iam.ListUsersInput{})

	p := iam.NewListUsersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_iam_user:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Users {
			fmt.Println(*r.UserName)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_iam_user: %s", err)
	}

}
