// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsTopicSubscription(client *Client) ([]Resource, error) {
	req := client.Snsconn.ListSubscriptionsRequest(&sns.ListSubscriptionsInput{})

	var result []Resource

	p := sns.NewListSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Subscriptions {

			result = append(result, Resource{
				Type:   "aws_sns_topic_subscription",
				ID:     *r.SubscriptionArn,
				Region: client.Snsconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
