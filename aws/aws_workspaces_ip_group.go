// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func ListWorkspacesIpGroup(client *Client) ([]Resource, error) {
	req := client.Workspacesconn.DescribeIpGroupsRequest(&workspaces.DescribeIpGroupsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Result) > 0 {
		for _, r := range resp.Result {

			result = append(result, Resource{
				Type:   "aws_workspaces_ip_group",
				ID:     *r.GroupId,
				Region: client.Workspacesconn.Config.Region,
			})
		}
	}

	return result, nil
}
