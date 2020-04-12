// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailStaticIp(client *Client) {
	req := client.lightsailconn.GetStaticIpsRequest(&lightsail.GetStaticIpsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_lightsail_static_ip: %s", err)
	} else {
		if len(resp.StaticIps) > 0 {
			fmt.Println("")
			fmt.Println("aws_lightsail_static_ip:")
			for _, r := range resp.StaticIps {
				fmt.Println(*r.Name)

			}
		}
	}

}
