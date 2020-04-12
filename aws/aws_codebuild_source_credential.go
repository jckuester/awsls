// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

func ListCodebuildSourceCredential(client *Client) error {
	req := client.codebuildconn.ListSourceCredentialsRequest(&codebuild.ListSourceCredentialsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.SourceCredentialsInfos) > 0 {
		for _, r := range resp.SourceCredentialsInfos {
			fmt.Println(*r.Arn)

		}
	}

	return nil
}
