// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
)

func ListKinesisAnalyticsApplication(client *Client) {
	req := client.kinesisanalyticsconn.ListApplicationsRequest(&kinesisanalytics.ListApplicationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_kinesis_analytics_application: %s", err)
	} else {
		if len(resp.ApplicationSummaries) > 0 {
			fmt.Println("")
			fmt.Println("aws_kinesis_analytics_application:")
			for _, r := range resp.ApplicationSummaries {
				fmt.Println(*r.ApplicationARN)

			}
		}
	}

}
