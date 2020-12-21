// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftGameSessionQueue(client *Client) ([]Resource, error) {
	req := client.Gameliftconn.DescribeGameSessionQueuesRequest(&gamelift.DescribeGameSessionQueuesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.GameSessionQueues) > 0 {

		for _, r := range resp.GameSessionQueues {

			result = append(result, Resource{
				Type:      "aws_gamelift_game_session_queue",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
