// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func ListSecretsmanagerSecret(client *Client) ([]Resource, error) {
	req := client.Secretsmanagerconn.ListSecretsRequest(&secretsmanager.ListSecretsInput{})

	var result []Resource

	p := secretsmanager.NewListSecretsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.SecretList {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:      "aws_secretsmanager_secret",
				ID:        *r.ARN,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
