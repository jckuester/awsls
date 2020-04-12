// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotThing(client *Client) {
	req := client.iotconn.ListThingsRequest(&iot.ListThingsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_iot_thing: %s", err)
	} else {
		if len(resp.Things) > 0 {
			fmt.Println("")
			fmt.Println("aws_iot_thing:")
			for _, r := range resp.Things {
				fmt.Println(*r.ThingName)

			}
		}
	}

}
