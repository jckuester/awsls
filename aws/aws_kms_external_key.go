// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func ListKmsExternalKey(client *Client) error {
	req := client.kmsconn.ListKeysRequest(&kms.ListKeysInput{})

	p := kms.NewListKeysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Keys {
			fmt.Println(*r.KeyId)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
