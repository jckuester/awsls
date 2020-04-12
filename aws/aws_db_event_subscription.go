// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbEventSubscription(client *Client) {
	req := client.rdsconn.DescribeEventSubscriptionsRequest(&rds.DescribeEventSubscriptionsInput{})

	p := rds.NewDescribeEventSubscriptionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_db_event_subscription:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSubscriptionsList {
			fmt.Println(*r.CustSubscriptionId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_db_event_subscription: %s", err)
	}

}
