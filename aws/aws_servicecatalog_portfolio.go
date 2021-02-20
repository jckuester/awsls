// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListServicecatalogPortfolio(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Servicecatalogconn.ListPortfoliosRequest(&servicecatalog.ListPortfoliosInput{})

	var result []terraform.Resource

	p := servicecatalog.NewListPortfoliosPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.PortfolioDetails {

			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_servicecatalog_portfolio",
				ID:        *r.Id,
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
