// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

func ListWorkspacesDirectory(client *Client) ([]Resource, error) {
	req := client.Workspacesconn.DescribeWorkspaceDirectoriesRequest(&workspaces.DescribeWorkspaceDirectoriesInput{})

	var result []Resource

	p := workspaces.NewDescribeWorkspaceDirectoriesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Directories {

			result = append(result, Resource{
				Type:      "aws_workspaces_directory",
				ID:        *r.DirectoryId,
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
