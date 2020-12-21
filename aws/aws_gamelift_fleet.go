// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftFleet(client *Client) ([]Resource, error) {
	req := client.Gameliftconn.ListFleetsRequest(&gamelift.ListFleetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.FleetIds) > 0 {

		for _, r := range resp.FleetIds {

			result = append(result, Resource{
				Type:      "aws_gamelift_fleet",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
