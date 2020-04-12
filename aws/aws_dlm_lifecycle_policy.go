// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

func ListDlmLifecyclePolicy(client *Client) {
	req := client.dlmconn.GetLifecyclePoliciesRequest(&dlm.GetLifecyclePoliciesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_dlm_lifecycle_policy: %s", err)
	} else {
		if len(resp.Policies) > 0 {
			fmt.Println("")
			fmt.Println("aws_dlm_lifecycle_policy:")
			for _, r := range resp.Policies {
				fmt.Println(*r.PolicyId)
				for k, v := range r.Tags {
					fmt.Printf("\t%s: %s\n", k, v)
				}
			}
		}
	}

}
