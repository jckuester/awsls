// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func ListCloudwatchEventBus(client *Client) ([]Resource, error) {
	req := client.Cloudwatcheventsconn.ListEventBusesRequest(&cloudwatchevents.ListEventBusesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.EventBuses) > 0 {

		for _, r := range resp.EventBuses {

			result = append(result, Resource{
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
