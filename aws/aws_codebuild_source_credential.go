// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

func ListCodebuildSourceCredential(client *Client) {
	req := client.codebuildconn.ListSourceCredentialsRequest(&codebuild.ListSourceCredentialsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_codebuild_source_credential: %s", err)
	} else {
		if len(resp.SourceCredentialsInfos) > 0 {
			fmt.Println("")
			fmt.Println("aws_codebuild_source_credential:")
			for _, r := range resp.SourceCredentialsInfos {
				fmt.Println(*r.Arn)

			}
		}
	}

}
