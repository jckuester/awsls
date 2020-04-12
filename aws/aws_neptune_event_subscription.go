// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
)

func ListNeptuneEventSubscription(client *Client) {
	req := client.neptuneconn.DescribeEventSubscriptionsRequest(&neptune.DescribeEventSubscriptionsInput{})

	p := neptune.NewDescribeEventSubscriptionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_neptune_event_subscription:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSubscriptionsList {
			fmt.Println(*r.CustSubscriptionId)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_neptune_event_subscription: %s", err)
	}

}
