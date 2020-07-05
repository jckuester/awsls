// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

func ListCodebuildSourceCredential(client *Client) ([]Resource, error) {
	req := client.Codebuildconn.ListSourceCredentialsRequest(&codebuild.ListSourceCredentialsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SourceCredentialsInfos) > 0 {
		for _, r := range resp.SourceCredentialsInfos {

			result = append(result, Resource{
				Type:   "aws_codebuild_source_credential",
				ID:     *r.Arn,
				Region: client.Codebuildconn.Config.Region,
			})
		}
	}

	return result, nil
}
