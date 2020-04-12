// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/worklink"
)

func ListWorklinkFleet(client *Client) {
	req := client.worklinkconn.ListFleetsRequest(&worklink.ListFleetsInput{})

	p := worklink.NewListFleetsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_worklink_fleet:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.FleetSummaryList {
			fmt.Println(*r.FleetArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_worklink_fleet: %s", err)
	}

}
