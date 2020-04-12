// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func ListMqConfiguration(client *Client) error {
	req := client.mqconn.ListConfigurationsRequest(&mq.ListConfigurationsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Configurations) > 0 {
		for _, r := range resp.Configurations {
			fmt.Println(*r.Id)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	return nil
}
