// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamServiceLinkedRole(client *Client) {
	req := client.iamconn.ListRolesRequest(&iam.ListRolesInput{})

	p := iam.NewListRolesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_iam_service_linked_role:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Roles {
			fmt.Println(*r.Arn)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_iam_service_linked_role: %s", err)
	}

}
