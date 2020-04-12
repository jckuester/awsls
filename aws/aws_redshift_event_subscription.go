// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftEventSubscription(client *Client) error {
	req := client.redshiftconn.DescribeEventSubscriptionsRequest(&redshift.DescribeEventSubscriptionsInput{})

	p := redshift.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSubscriptionsList {
			fmt.Println(*r.CustSubscriptionId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
