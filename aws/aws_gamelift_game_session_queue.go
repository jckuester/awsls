// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGameliftGameSessionQueue(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Gameliftconn.DescribeGameSessionQueuesRequest(&gamelift.DescribeGameSessionQueuesInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.GameSessionQueues) > 0 {

		for _, r := range resp.GameSessionQueues {

			result = append(result, terraform.Resource{
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
