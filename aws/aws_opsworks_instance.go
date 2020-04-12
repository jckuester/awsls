// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksInstance(client *Client) {
	req := client.opsworksconn.DescribeInstancesRequest(&opsworks.DescribeInstancesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_opsworks_instance: %s", err)
	} else {
		if len(resp.Instances) > 0 {
			fmt.Println("")
			fmt.Println("aws_opsworks_instance:")
			for _, r := range resp.Instances {
				fmt.Println(*r.InstanceId)

			}
		}
	}

}
