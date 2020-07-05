// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
)

func ListGlobalacceleratorAccelerator(client *Client) ([]Resource, error) {
	req := client.Globalacceleratorconn.ListAcceleratorsRequest(&globalaccelerator.ListAcceleratorsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Accelerators) > 0 {
		for _, r := range resp.Accelerators {

			t := *r.CreatedTime
			result = append(result, Resource{
				Type:   "aws_globalaccelerator_accelerator",
				ID:     *r.AcceleratorArn,
				Region: client.Globalacceleratorconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
