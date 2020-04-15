// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailInstance(client *Client) ([]Resource, error) {
	req := client.lightsailconn.GetInstancesRequest(&lightsail.GetInstancesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Instances) > 0 {
		for _, r := range resp.Instances {
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type: "aws_lightsail_instance",
				ID:   *r.Name,
				Tags: tags,
			})
		}
	}

	return result, nil
}
