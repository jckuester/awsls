// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func ListWorkspacesWorkspace(client *Client) ([]Resource, error) {
	req := client.Workspacesconn.DescribeWorkspacesRequest(&workspaces.DescribeWorkspacesInput{})

	var result []Resource

	p := workspaces.NewDescribeWorkspacesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Workspaces {

			result = append(result, Resource{
				Type:      "aws_workspaces_workspace",
				ID:        *r.WorkspaceId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
