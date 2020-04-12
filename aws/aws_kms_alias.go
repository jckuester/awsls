// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func ListKmsAlias(client *Client) error {
	req := client.kmsconn.ListAliasesRequest(&kms.ListAliasesInput{})

	p := kms.NewListAliasesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Aliases {
			fmt.Println(*r.AliasName)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
