// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudwatchEventBus(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Cloudwatcheventsconn.ListEventBuses(ctx, &cloudwatchevents.ListEventBusesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.EventBuses) > 0 {

		for _, r := range resp.EventBuses {

			result = append(result, terraform.Resource{
				Type:      "aws_cloudwatch_event_bus",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
