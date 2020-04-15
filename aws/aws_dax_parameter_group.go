// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func ListDaxParameterGroup(client *Client) ([]Resource, error) {
	req := client.daxconn.DescribeParameterGroupsRequest(&dax.DescribeParameterGroupsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ParameterGroups) > 0 {
		for _, r := range resp.ParameterGroups {

			result = append(result, Resource{
				Type: "aws_dax_parameter_group",
				ID:   *r.ParameterGroupName,
			})
		}
	}

	return result, nil
}
