// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftGameSessionQueue(client *Client) {
	req := client.gameliftconn.DescribeGameSessionQueuesRequest(&gamelift.DescribeGameSessionQueuesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_gamelift_game_session_queue: %s", err)
	} else {
		if len(resp.GameSessionQueues) > 0 {
			fmt.Println("")
			fmt.Println("aws_gamelift_game_session_queue:")
			for _, r := range resp.GameSessionQueues {
				fmt.Println(*r.Name)

			}
		}
	}

}
