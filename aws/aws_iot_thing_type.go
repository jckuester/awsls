// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotThingType(client *Client) ([]Resource, error) {
	req := client.Iotconn.ListThingTypesRequest(&iot.ListThingTypesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ThingTypes) > 0 {
		for _, r := range resp.ThingTypes {

			result = append(result, Resource{
				Type:   "aws_iot_thing_type",
				ID:     *r.ThingTypeName,
				Region: client.Iotconn.Config.Region,
			})
		}
	}

	return result, nil
}
