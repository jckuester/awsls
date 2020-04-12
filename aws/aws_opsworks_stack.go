// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksStack(client *Client) error {
	req := client.opsworksconn.DescribeStacksRequest(&opsworks.DescribeStacksInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Stacks) > 0 {
		for _, r := range resp.Stacks {
			fmt.Println(*r.StackId)
		}
	}

	return nil
}
