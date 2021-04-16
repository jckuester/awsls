// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListKinesisanalyticsv2Application(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Kinesisanalyticsv2conn.ListApplications(ctx, &kinesisanalyticsv2.ListApplicationsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ApplicationSummaries) > 0 {

		for _, r := range resp.ApplicationSummaries {

			result = append(result, terraform.Resource{
				Type:      "aws_kinesisanalyticsv2_application",
				ID:        *r.ApplicationARN,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
