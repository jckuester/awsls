// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func ListMskConfiguration(client *Client) error {
	req := client.kafkaconn.ListConfigurationsRequest(&kafka.ListConfigurationsInput{})

	p := kafka.NewListConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Configurations {
			fmt.Println(*r.Arn)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
