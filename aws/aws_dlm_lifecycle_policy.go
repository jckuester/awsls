// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

func ListDlmLifecyclePolicy(client *Client) ([]Resource, error) {
	req := client.dlmconn.GetLifecyclePoliciesRequest(&dlm.GetLifecyclePoliciesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Policies) > 0 {
		for _, r := range resp.Policies {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type: "aws_dlm_lifecycle_policy",
				ID:   *r.PolicyId,
				Tags: tags,
			})
		}
	}

	return result, nil
}
