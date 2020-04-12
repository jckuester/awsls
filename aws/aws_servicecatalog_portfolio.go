// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

func ListServicecatalogPortfolio(client *Client) {
	req := client.servicecatalogconn.ListPortfoliosRequest(&servicecatalog.ListPortfoliosInput{})

	p := servicecatalog.NewListPortfoliosPaginator(req)
	fmt.Println("")
	fmt.Println("aws_servicecatalog_portfolio:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.PortfolioDetails {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_servicecatalog_portfolio: %s", err)
	}

}
