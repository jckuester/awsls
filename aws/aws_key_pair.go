// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListKeyPair(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeKeyPairsRequest(&ec2.DescribeKeyPairsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.KeyPairs) > 0 {
		for _, r := range resp.KeyPairs {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:    "aws_key_pair",
				ID:      *r.KeyName,
				Profile: client.Profile,
				Region:  client.Region,
				Tags:    tags,
			})
		}
	}

	return result, nil
}
