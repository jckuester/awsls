// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailKeyPair(client *Client) ([]Resource, error) {
	req := client.lightsailconn.GetKeyPairsRequest(&lightsail.GetKeyPairsInput{})

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
				Type: "aws_lightsail_key_pair",
				ID:   *r.Name,
				Tags: tags,
			})
		}
	}

	return result, nil
}
