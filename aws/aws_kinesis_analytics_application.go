// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
)

func ListKinesisAnalyticsApplication(client *Client) ([]Resource, error) {
	req := client.kinesisanalyticsconn.ListApplicationsRequest(&kinesisanalytics.ListApplicationsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ApplicationSummaries) > 0 {
		for _, r := range resp.ApplicationSummaries {

			result = append(result, Resource{
				Type: "aws_kinesis_analytics_application",
				ID:   *r.ApplicationARN,
			})
		}
	}

	return result, nil
}
