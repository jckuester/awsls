// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftGameSessionQueue(client *Client) error {
	req := client.gameliftconn.DescribeGameSessionQueuesRequest(&gamelift.DescribeGameSessionQueuesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.GameSessionQueues) > 0 {
		for _, r := range resp.GameSessionQueues {
			fmt.Println(*r.Name)
		}
	}

	return nil
}
