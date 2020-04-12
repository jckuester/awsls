// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func ListMskConfiguration(client *Client) {
	req := client.kafkaconn.ListConfigurationsRequest(&kafka.ListConfigurationsInput{})

	p := kafka.NewListConfigurationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_msk_configuration:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Configurations {
			fmt.Println(*r.Arn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_msk_configuration: %s", err)
	}

}
