// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func ListKmsKey(client *Client) {
	req := client.kmsconn.ListKeysRequest(&kms.ListKeysInput{})

	p := kms.NewListKeysPaginator(req)
	fmt.Println("")
	fmt.Println("aws_kms_key:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Keys {
			fmt.Println(*r.KeyId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_kms_key: %s", err)
	}

}
