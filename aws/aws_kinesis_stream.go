// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListKinesisStream(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Kinesisconn.ListStreams(ctx, &kinesis.ListStreamsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.StreamNames) > 0 {

		for _, r := range resp.StreamNames {

			result = append(result, terraform.Resource{
				Type:      "aws_kinesis_stream",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
