// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

func ListKinesisStream(client *Client) ([]Resource, error) {
	req := client.Kinesisconn.ListStreamsRequest(&kinesis.ListStreamsInput{})

	var result []Resource

	p := kinesis.NewListStreamsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.StreamNames {

			result = append(result, Resource{
				Type:      "aws_kinesis_stream",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
