// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamAccessKey(client *Client) {
	req := client.iamconn.ListAccessKeysRequest(&iam.ListAccessKeysInput{})

	p := iam.NewListAccessKeysPaginator(req)
	fmt.Println("")
	fmt.Println("aws_iam_access_key:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.AccessKeyMetadata {
			fmt.Println(*r.AccessKeyId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_iam_access_key: %s", err)
	}

}
