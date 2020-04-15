// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotThing(client *Client) ([]Resource, error) {
	req := client.iotconn.ListThingsRequest(&iot.ListThingsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Things) > 0 {
		for _, r := range resp.Things {

			result = append(result, Resource{
				Type: "aws_iot_thing",
				ID:   *r.ThingName,
			})
		}
	}

	return result, nil
}
