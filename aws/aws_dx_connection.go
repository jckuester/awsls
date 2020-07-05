// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxConnection(client *Client) ([]Resource, error) {
	req := client.Directconnectconn.DescribeConnectionsRequest(&directconnect.DescribeConnectionsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Connections) > 0 {
		for _, r := range resp.Connections {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_dx_connection",
				ID:     *r.ConnectionId,
				Region: client.Directconnectconn.Config.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
