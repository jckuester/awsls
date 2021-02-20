// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDbEventSubscription(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Rdsconn.DescribeEventSubscriptionsRequest(&rds.DescribeEventSubscriptionsInput{})

	var result []terraform.Resource

	p := rds.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.EventSubscriptionsList {

			result = append(result, terraform.Resource{
				Type:      "aws_db_event_subscription",
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
