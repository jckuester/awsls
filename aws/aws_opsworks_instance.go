// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

func ListOpsworksInstance(client *Client) ([]Resource, error) {
	req := client.opsworksconn.DescribeInstancesRequest(&opsworks.DescribeInstancesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Instances) > 0 {
		for _, r := range resp.Instances {

			result = append(result, Resource{
				Type:   "aws_opsworks_instance",
				ID:     *r.InstanceId,
				Region: client.opsworksconn.Config.Region,
			})
		}
	}

	return result, nil
}
