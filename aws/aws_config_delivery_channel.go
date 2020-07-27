// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigDeliveryChannel(client *Client) ([]Resource, error) {
	req := client.Configserviceconn.DescribeDeliveryChannelsRequest(&configservice.DescribeDeliveryChannelsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.DeliveryChannels) > 0 {
		for _, r := range resp.DeliveryChannels {

			result = append(result, Resource{
				Type:      "aws_config_delivery_channel",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
