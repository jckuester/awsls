// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmAssociation(client *Client) error {
	req := client.ssmconn.ListAssociationsRequest(&ssm.ListAssociationsInput{})

	p := ssm.NewListAssociationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Associations {
			fmt.Println(*r.AssociationId)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
