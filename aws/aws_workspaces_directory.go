// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWorkspacesDirectory(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := workspaces.NewDescribeWorkspaceDirectoriesPaginator(client.Workspacesconn, &workspaces.DescribeWorkspaceDirectoriesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Directories {

			result = append(result, terraform.Resource{
				Type:      "aws_workspaces_directory",
				ID:        *r.DirectoryId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
