// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListKinesisAnalyticsApplication(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Kinesisanalyticsconn.ListApplications(ctx, &kinesisanalytics.ListApplicationsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ApplicationSummaries) > 0 {

		for _, r := range resp.ApplicationSummaries {

			result = append(result, terraform.Resource{
				Type:      "aws_kinesis_analytics_application",
				ID:        *r.ApplicationARN,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
