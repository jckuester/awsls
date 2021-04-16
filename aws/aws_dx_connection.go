// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDxConnection(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Directconnectconn.DescribeConnections(ctx, &directconnect.DescribeConnectionsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.Connections) > 0 {

		for _, r := range resp.Connections {

			result = append(result, terraform.Resource{
				Type:      "aws_dx_connection",
				ID:        *r.ConnectionId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
