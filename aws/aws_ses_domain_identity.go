// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSesDomainIdentity(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Sesconn.ListIdentitiesRequest(&ses.ListIdentitiesInput{})

	var result []terraform.Resource

	p := ses.NewListIdentitiesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Identities {

			result = append(result, terraform.Resource{
				Type:      "aws_ses_domain_identity",
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
