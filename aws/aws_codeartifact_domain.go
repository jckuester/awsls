// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodeartifactDomain(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codeartifactconn.ListDomainsRequest(&codeartifact.ListDomainsInput{})

	var result []terraform.Resource

	p := codeartifact.NewListDomainsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Domains {

			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_codeartifact_domain",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
