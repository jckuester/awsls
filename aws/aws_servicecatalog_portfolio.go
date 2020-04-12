// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

func ListServicecatalogPortfolio(client *Client) error {
	req := client.servicecatalogconn.ListPortfoliosRequest(&servicecatalog.ListPortfoliosInput{})

	p := servicecatalog.NewListPortfoliosPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.PortfolioDetails {
			fmt.Println(*r.Id)

			fmt.Printf("CreatedAt: %s\n", *r.CreatedTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
