// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListNeptuneEventSubscription(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := neptune.NewDescribeEventSubscriptionsPaginator(client.Neptuneconn, &neptune.DescribeEventSubscriptionsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

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

	return result, nil
}
