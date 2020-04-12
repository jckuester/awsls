// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsTopicSubscription(client *Client) error {
	req := client.snsconn.ListSubscriptionsRequest(&sns.ListSubscriptionsInput{})

	p := sns.NewListSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Subscriptions {
			fmt.Println(*r.SubscriptionArn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
