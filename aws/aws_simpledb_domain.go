// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/simpledb"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSimpledbDomain(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Simpledbconn.ListDomainsRequest(&simpledb.ListDomainsInput{})

	var result []terraform.Resource

	p := simpledb.NewListDomainsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.DomainNames {

			result = append(result, terraform.Resource{
				Type:      "aws_simpledb_domain",
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
