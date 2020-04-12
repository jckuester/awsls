// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamPolicy(client *Client) {
	req := client.iamconn.ListPoliciesRequest(&iam.ListPoliciesInput{})

	p := iam.NewListPoliciesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_iam_policy:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Policies {
			fmt.Println(*r.Arn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_iam_policy: %s", err)
	}

}
