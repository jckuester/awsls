// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func ListMqConfiguration(client *Client) {
	req := client.mqconn.ListConfigurationsRequest(&mq.ListConfigurationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_mq_configuration: %s", err)
	} else {
		if len(resp.Configurations) > 0 {
			fmt.Println("")
			fmt.Println("aws_mq_configuration:")
			for _, r := range resp.Configurations {
				fmt.Println(*r.Id)
				for k, v := range r.Tags {
					fmt.Printf("\t%s: %s\n", k, v)
				}
			}
		}
	}

}
