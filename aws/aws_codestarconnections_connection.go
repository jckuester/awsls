// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodestarconnectionsConnection(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codestarconnectionsconn.ListConnectionsRequest(&codestarconnections.ListConnectionsInput{})

	var result []terraform.Resource

	p := codestarconnections.NewListConnectionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Connections {

			result = append(result, terraform.Resource{
				Type:      "aws_codestarconnections_connection",
				ID:        *r.ConnectionArn,
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
