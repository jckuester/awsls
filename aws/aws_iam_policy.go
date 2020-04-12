// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamPolicy(client *Client) error {
	req := client.iamconn.ListPoliciesRequest(&iam.ListPoliciesInput{})

	p := iam.NewListPoliciesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Policies {
			fmt.Println(*r.Arn)

			fmt.Printf("CreatedAt: %s\n", *r.CreateDate)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
