// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailStaticIp(client *Client) error {
	req := client.lightsailconn.GetStaticIpsRequest(&lightsail.GetStaticIpsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.StaticIps) > 0 {
		for _, r := range resp.StaticIps {
			fmt.Println(*r.Name)

		}
	}

	return nil
}
