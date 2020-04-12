// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

func ListDlmLifecyclePolicy(client *Client) error {
	req := client.dlmconn.GetLifecyclePoliciesRequest(&dlm.GetLifecyclePoliciesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Policies) > 0 {
		for _, r := range resp.Policies {
			fmt.Println(*r.PolicyId)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	return nil
}
