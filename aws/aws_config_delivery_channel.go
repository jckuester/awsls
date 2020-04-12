// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigDeliveryChannel(client *Client) error {
	req := client.configserviceconn.DescribeDeliveryChannelsRequest(&configservice.DescribeDeliveryChannelsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.DeliveryChannels) > 0 {
		for _, r := range resp.DeliveryChannels {
			fmt.Println(*r.Name)
		}
	}

	return nil
}
