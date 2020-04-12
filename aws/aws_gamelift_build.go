// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftBuild(client *Client) {
	req := client.gameliftconn.ListBuildsRequest(&gamelift.ListBuildsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_gamelift_build: %s", err)
	} else {
		if len(resp.Builds) > 0 {
			fmt.Println("")
			fmt.Println("aws_gamelift_build:")
			for _, r := range resp.Builds {
				fmt.Println(*r.BuildId)

			}
		}
	}

}
