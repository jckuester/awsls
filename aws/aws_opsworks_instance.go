// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksInstance(client *Client) error {
	req := client.opsworksconn.DescribeInstancesRequest(&opsworks.DescribeInstancesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Instances) > 0 {
		for _, r := range resp.Instances {
			fmt.Println(*r.InstanceId)

		}
	}

	return nil
}
