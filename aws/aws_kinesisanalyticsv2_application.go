// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
)

func ListKinesisanalyticsv2Application(client *Client) ([]Resource, error) {
	req := client.Kinesisanalyticsv2conn.ListApplicationsRequest(&kinesisanalyticsv2.ListApplicationsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ApplicationSummaries) > 0 {

		for _, r := range resp.ApplicationSummaries {

			result = append(result, Resource{
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
