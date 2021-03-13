// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListKinesisFirehoseDeliveryStream(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Firehoseconn.ListDeliveryStreams(ctx, &firehose.ListDeliveryStreamsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.DeliveryStreamNames) > 0 {

		for _, r := range resp.DeliveryStreamNames {

			result = append(result, terraform.Resource{
				Type:      "aws_kinesis_firehose_delivery_stream",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
