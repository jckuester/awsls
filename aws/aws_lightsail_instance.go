// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailInstance(client *Client) {
	req := client.lightsailconn.GetInstancesRequest(&lightsail.GetInstancesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_lightsail_instance: %s", err)
	} else {
		if len(resp.Instances) > 0 {
			fmt.Println("")
			fmt.Println("aws_lightsail_instance:")
			for _, r := range resp.Instances {
				fmt.Println(*r.Name)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}
