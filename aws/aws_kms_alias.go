// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func ListKmsAlias(client *Client) ([]Resource, error) {
	req := client.kmsconn.ListAliasesRequest(&kms.ListAliasesInput{})

	var result []Resource

	p := kms.NewListAliasesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Aliases {

			result = append(result, Resource{
				Type: "aws_kms_alias",
				ID:   *r.AliasName,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
