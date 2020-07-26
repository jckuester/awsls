// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksStack(client *Client) ([]Resource, error) {
	req := client.Opsworksconn.DescribeStacksRequest(&opsworks.DescribeStacksInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Stacks) > 0 {
		for _, r := range resp.Stacks {

			result = append(result, Resource{
				Type:    "aws_opsworks_stack",
				ID:      *r.StackId,
				Profile: client.Profile,
				Region:  client.Region,
			})
		}
	}

	return result, nil
}
