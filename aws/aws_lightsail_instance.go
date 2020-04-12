// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

func ListLightsailInstance(client *Client) error {
	req := client.lightsailconn.GetInstancesRequest(&lightsail.GetInstancesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Instances) > 0 {
		for _, r := range resp.Instances {
			fmt.Println(*r.Name)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	return nil
}
