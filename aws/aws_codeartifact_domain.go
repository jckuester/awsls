// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
)

func ListCodeartifactDomain(client *Client) ([]Resource, error) {
	req := client.Codeartifactconn.ListDomainsRequest(&codeartifact.ListDomainsInput{})

	var result []Resource

	p := codeartifact.NewListDomainsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Domains {

			t := *r.CreatedTime
			result = append(result, Resource{
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
