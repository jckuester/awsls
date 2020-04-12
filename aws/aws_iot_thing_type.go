// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotThingType(client *Client) error {
	req := client.iotconn.ListThingTypesRequest(&iot.ListThingTypesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.ThingTypes) > 0 {
		for _, r := range resp.ThingTypes {
			fmt.Println(*r.ThingTypeName)
		}
	}

	return nil
}
