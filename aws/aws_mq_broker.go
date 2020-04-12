// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func ListMqBroker(client *Client) {
	req := client.mqconn.ListBrokersRequest(&mq.ListBrokersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_mq_broker: %s", err)
	} else {
		if len(resp.BrokerSummaries) > 0 {
			fmt.Println("")
			fmt.Println("aws_mq_broker:")
			for _, r := range resp.BrokerSummaries {
				fmt.Println(*r.BrokerId)

			}
		}
	}

}
