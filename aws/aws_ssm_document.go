// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmDocument(client *Client) {
	req := client.ssmconn.ListDocumentsRequest(&ssm.ListDocumentsInput{})

	p := ssm.NewListDocumentsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_ssm_document:")
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
		log.Printf("aws_ssm_document: %s", err)
	}

}
