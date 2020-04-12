// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigDeliveryChannel(client *Client) {
	req := client.configserviceconn.DescribeDeliveryChannelsRequest(&configservice.DescribeDeliveryChannelsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_config_delivery_channel: %s", err)
	} else {
		if len(resp.DeliveryChannels) > 0 {
			fmt.Println("")
			fmt.Println("aws_config_delivery_channel:")
			for _, r := range resp.DeliveryChannels {
				fmt.Println(*r.Name)

			}
		}
	}

}
