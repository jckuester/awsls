// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigDeliveryChannel(client *Client) ([]Resource, error) {
	req := client.configserviceconn.DescribeDeliveryChannelsRequest(&configservice.DescribeDeliveryChannelsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.DeliveryChannels) > 0 {
		for _, r := range resp.DeliveryChannels {

			result = append(result, Resource{
				Type:   "aws_config_delivery_channel",
				ID:     *r.Name,
				Region: client.configserviceconn.Config.Region,
			})
		}
	}

	return result, nil
}
