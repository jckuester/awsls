// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

func ListCodebuildProject(client *Client) ([]Resource, error) {
	req := client.Codebuildconn.ListProjectsRequest(&codebuild.ListProjectsInput{})

	var result []Resource

	p := codebuild.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Projects {

			result = append(result, Resource{
				Type:      "aws_codebuild_project",
				ID:        r,
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
