// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListNeptuneEventSubscription(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Neptuneconn.DescribeEventSubscriptionsRequest(&neptune.DescribeEventSubscriptionsInput{})

	var result []terraform.Resource

	p := neptune.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.EventSubscriptionsList {

			result = append(result, terraform.Resource{
				Type:      "aws_neptune_event_subscription",
				ID:        *r.CustSubscriptionId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
