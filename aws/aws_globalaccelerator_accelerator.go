// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
)

func ListGlobalacceleratorAccelerator(client *Client) error {
	req := client.globalacceleratorconn.ListAcceleratorsRequest(&globalaccelerator.ListAcceleratorsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Accelerators) > 0 {
		for _, r := range resp.Accelerators {
			fmt.Println(*r.AcceleratorArn)

			fmt.Printf("CreatedAt: %s\n", *r.CreatedTime)
		}
	}

	return nil
}
