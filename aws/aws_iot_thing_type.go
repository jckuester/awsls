// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotThingType(client *Client) {
	req := client.iotconn.ListThingTypesRequest(&iot.ListThingTypesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_iot_thing_type: %s", err)
	} else {
		if len(resp.ThingTypes) > 0 {
			fmt.Println("")
			fmt.Println("aws_iot_thing_type:")
			for _, r := range resp.ThingTypes {
				fmt.Println(*r.ThingTypeName)

			}
		}
	}

}
