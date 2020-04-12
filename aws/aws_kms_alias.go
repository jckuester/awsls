// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func ListKmsAlias(client *Client) {
	req := client.kmsconn.ListAliasesRequest(&kms.ListAliasesInput{})

	p := kms.NewListAliasesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_kms_alias:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Aliases {
			fmt.Println(*r.AliasName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_kms_alias: %s", err)
	}

}
