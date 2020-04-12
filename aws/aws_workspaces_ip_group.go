// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func ListWorkspacesIpGroup(client *Client) error {
	req := client.workspacesconn.DescribeIpGroupsRequest(&workspaces.DescribeIpGroupsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Result) > 0 {
		for _, r := range resp.Result {
			fmt.Println(*r.GroupId)

		}
	}

	return nil
}
