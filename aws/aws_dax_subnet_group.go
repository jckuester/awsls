// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func ListDaxSubnetGroup(client *Client) ([]Resource, error) {
	req := client.daxconn.DescribeSubnetGroupsRequest(&dax.DescribeSubnetGroupsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SubnetGroups) > 0 {
		for _, r := range resp.SubnetGroups {

			result = append(result, Resource{
				Type: "aws_dax_subnet_group",
				ID:   *r.SubnetGroupName,
			})
		}
	}

	return result, nil
}
