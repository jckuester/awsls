// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotThing(client *Client) error {
	req := client.iotconn.ListThingsRequest(&iot.ListThingsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Things) > 0 {
		for _, r := range resp.Things {
			fmt.Println(*r.ThingName)
		}
	}

	return nil
}
