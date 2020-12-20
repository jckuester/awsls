// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamAccountAlias(client *Client) ([]Resource, error) {
	req := client.Iamconn.ListAccountAliasesRequest(&iam.ListAccountAliasesInput{})

	var result []Resource

	p := iam.NewListAccountAliasesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.AccountAliases {

			result = append(result, Resource{
				Type:      "aws_iam_account_alias",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
