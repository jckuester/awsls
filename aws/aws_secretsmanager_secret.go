// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func ListSecretsmanagerSecret(client *Client) {
	req := client.secretsmanagerconn.ListSecretsRequest(&secretsmanager.ListSecretsInput{})

	p := secretsmanager.NewListSecretsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_secretsmanager_secret:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecretList {
			fmt.Println(*r.ARN)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_secretsmanager_secret: %s", err)
	}

}
