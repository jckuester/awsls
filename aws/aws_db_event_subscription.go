// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbEventSubscription(client *Client) error {
	req := client.rdsconn.DescribeEventSubscriptionsRequest(&rds.DescribeEventSubscriptionsInput{})

	p := rds.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSubscriptionsList {
			fmt.Println(*r.CustSubscriptionId)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
