// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamGroup(client *Client) {
	req := client.iamconn.ListGroupsRequest(&iam.ListGroupsInput{})

	p := iam.NewListGroupsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_iam_group:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Groups {
			fmt.Println(*r.GroupName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_iam_group: %s", err)
	}

}
