// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func ListMskConfiguration(client *Client) ([]Resource, error) {
	req := client.Kafkaconn.ListConfigurationsRequest(&kafka.ListConfigurationsInput{})

	var result []Resource

	p := kafka.NewListConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Configurations {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:   "aws_msk_configuration",
				ID:     *r.Arn,
				Region: client.Kafkaconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
