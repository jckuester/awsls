// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotPolicy(client *Client) error {
	req := client.iotconn.ListPoliciesRequest(&iot.ListPoliciesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Policies) > 0 {
		for _, r := range resp.Policies {
			fmt.Println(*r.PolicyName)
		}
	}

	return nil
}
