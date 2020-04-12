// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsTopicSubscription(client *Client) {
	req := client.snsconn.ListSubscriptionsRequest(&sns.ListSubscriptionsInput{})

	p := sns.NewListSubscriptionsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_sns_topic_subscription:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Subscriptions {
			fmt.Println(*r.SubscriptionArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_sns_topic_subscription: %s", err)
	}

}
