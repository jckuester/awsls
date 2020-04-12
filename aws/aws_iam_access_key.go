// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamAccessKey(client *Client) error {
	req := client.iamconn.ListAccessKeysRequest(&iam.ListAccessKeysInput{})

	p := iam.NewListAccessKeysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.AccessKeyMetadata {
			fmt.Println(*r.AccessKeyId)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
