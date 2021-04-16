// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodestarnotificationsNotificationRule(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := codestarnotifications.NewListNotificationRulesPaginator(client.Codestarnotificationsconn, &codestarnotifications.ListNotificationRulesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

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

	return result, nil
}
