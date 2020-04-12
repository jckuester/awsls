// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamGroup(client *Client) error {
	req := client.iamconn.ListGroupsRequest(&iam.ListGroupsInput{})

	p := iam.NewListGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Groups {
			fmt.Println(*r.GroupName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
