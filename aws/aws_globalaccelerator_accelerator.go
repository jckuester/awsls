// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
)

func ListGlobalacceleratorAccelerator(client *Client) {
	req := client.globalacceleratorconn.ListAcceleratorsRequest(&globalaccelerator.ListAcceleratorsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_globalaccelerator_accelerator: %s", err)
	} else {
		if len(resp.Accelerators) > 0 {
			fmt.Println("")
			fmt.Println("aws_globalaccelerator_accelerator:")
			for _, r := range resp.Accelerators {
				fmt.Println(*r.AcceleratorArn)

			}
		}
	}

}
