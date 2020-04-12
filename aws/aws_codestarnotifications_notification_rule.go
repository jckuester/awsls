// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
)

func ListCodestarnotificationsNotificationRule(client *Client) error {
	req := client.codestarnotificationsconn.ListNotificationRulesRequest(&codestarnotifications.ListNotificationRulesInput{})

	p := codestarnotifications.NewListNotificationRulesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NotificationRules {
			fmt.Println(*r.Arn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
