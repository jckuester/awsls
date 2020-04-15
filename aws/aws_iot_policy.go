// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotPolicy(client *Client) ([]Resource, error) {
	req := client.iotconn.ListPoliciesRequest(&iot.ListPoliciesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Policies) > 0 {
		for _, r := range resp.Policies {

			result = append(result, Resource{
				Type: "aws_iot_policy",
				ID:   *r.PolicyName,
			})
		}
	}

	return result, nil
}
