// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/worklink"
)

func ListWorklinkFleet(client *Client) error {
	req := client.worklinkconn.ListFleetsRequest(&worklink.ListFleetsInput{})

	p := worklink.NewListFleetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.FleetSummaryList {
			fmt.Println(*r.FleetArn)

			fmt.Printf("CreatedAt: %s\n", *r.CreatedTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
