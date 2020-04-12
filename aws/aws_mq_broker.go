// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func ListMqBroker(client *Client) error {
	req := client.mqconn.ListBrokersRequest(&mq.ListBrokersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.BrokerSummaries) > 0 {
		for _, r := range resp.BrokerSummaries {
			fmt.Println(*r.BrokerId)

		}
	}

	return nil
}
