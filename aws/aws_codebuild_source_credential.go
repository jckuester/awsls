// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodebuildSourceCredential(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codebuildconn.ListSourceCredentialsRequest(&codebuild.ListSourceCredentialsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SourceCredentialsInfos) > 0 {

		for _, r := range resp.SourceCredentialsInfos {

			result = append(result, terraform.Resource{
				Type:      "aws_codebuild_source_credential",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
