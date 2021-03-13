// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListConfigDeliveryChannel(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Configserviceconn.DescribeDeliveryChannels(ctx, &configservice.DescribeDeliveryChannelsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.DeliveryChannels) > 0 {

		for _, r := range resp.DeliveryChannels {

			result = append(result, terraform.Resource{
				Type:      "aws_config_delivery_channel",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
