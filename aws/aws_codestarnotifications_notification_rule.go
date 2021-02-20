// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodestarnotificationsNotificationRule(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codestarnotificationsconn.ListNotificationRulesRequest(&codestarnotifications.ListNotificationRulesInput{})

	var result []terraform.Resource

	p := codestarnotifications.NewListNotificationRulesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.NotificationRules {

			result = append(result, terraform.Resource{
				Type:      "aws_codestarnotifications_notification_rule",
				ID:        *r.Arn,
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
