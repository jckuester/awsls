// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func ListDxLag(client *Client) ([]Resource, error) {
	req := client.Directconnectconn.DescribeLagsRequest(&directconnect.DescribeLagsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Lags) > 0 {
		for _, r := range resp.Lags {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_dx_lag",
				ID:     *r.LagId,
				Region: client.Directconnectconn.Config.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
