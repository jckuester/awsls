// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksStack(client *Client) {
	req := client.opsworksconn.DescribeStacksRequest(&opsworks.DescribeStacksInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_opsworks_stack: %s", err)
	} else {
		if len(resp.Stacks) > 0 {
			fmt.Println("")
			fmt.Println("aws_opsworks_stack:")
			for _, r := range resp.Stacks {
				fmt.Println(*r.StackId)

			}
		}
	}

}
