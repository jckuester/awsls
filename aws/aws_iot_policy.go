// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotPolicy(client *Client) {
	req := client.iotconn.ListPoliciesRequest(&iot.ListPoliciesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_iot_policy: %s", err)
	} else {
		if len(resp.Policies) > 0 {
			fmt.Println("")
			fmt.Println("aws_iot_policy:")
			for _, r := range resp.Policies {
				fmt.Println(*r.PolicyName)

			}
		}
	}

}
