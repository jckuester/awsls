// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
)

func ListKinesisAnalyticsApplication(client *Client) error {
	req := client.kinesisanalyticsconn.ListApplicationsRequest(&kinesisanalytics.ListApplicationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.ApplicationSummaries) > 0 {
		for _, r := range resp.ApplicationSummaries {
			fmt.Println(*r.ApplicationARN)

		}
	}

	return nil
}
