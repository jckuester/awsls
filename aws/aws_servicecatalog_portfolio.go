// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

func ListServicecatalogPortfolio(client *Client) ([]Resource, error) {
	req := client.Servicecatalogconn.ListPortfoliosRequest(&servicecatalog.ListPortfoliosInput{})

	var result []Resource

	p := servicecatalog.NewListPortfoliosPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.PortfolioDetails {

			t := *r.CreatedTime
			result = append(result, Resource{
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
