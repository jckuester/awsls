// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamServiceLinkedRole(client *Client) error {
	req := client.iamconn.ListRolesRequest(&iam.ListRolesInput{})

	p := iam.NewListRolesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Roles {
			fmt.Println(*r.Arn)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
			fmt.Printf("CreatedAt: %s\n", *r.CreateDate)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
