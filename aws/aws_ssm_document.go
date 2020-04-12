// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmDocument(client *Client) error {
	req := client.ssmconn.ListDocumentsRequest(&ssm.ListDocumentsInput{})

	p := ssm.NewListDocumentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DocumentIdentifiers {
			fmt.Println(*r.Name)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
