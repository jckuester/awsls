// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func ListDaxParameterGroup(client *Client) error {
	req := client.daxconn.DescribeParameterGroupsRequest(&dax.DescribeParameterGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.ParameterGroups) > 0 {
		for _, r := range resp.ParameterGroups {
			fmt.Println(*r.ParameterGroupName)
		}
	}

	return nil
}
