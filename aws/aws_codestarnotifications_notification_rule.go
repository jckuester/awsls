// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
)

func ListCodestarnotificationsNotificationRule(client *Client) {
	req := client.codestarnotificationsconn.ListNotificationRulesRequest(&codestarnotifications.ListNotificationRulesInput{})

	p := codestarnotifications.NewListNotificationRulesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_codestarnotifications_notification_rule:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NotificationRules {
			fmt.Println(*r.Arn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_codestarnotifications_notification_rule: %s", err)
	}

}
