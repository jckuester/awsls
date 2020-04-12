// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dax"
)

func ListDaxSubnetGroup(client *Client) error {
	req := client.daxconn.DescribeSubnetGroupsRequest(&dax.DescribeSubnetGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.SubnetGroups) > 0 {
		for _, r := range resp.SubnetGroups {
			fmt.Println(*r.SubnetGroupName)

		}
	}

	return nil
}
