// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftBuild(client *Client) error {
	req := client.gameliftconn.ListBuildsRequest(&gamelift.ListBuildsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Builds) > 0 {
		for _, r := range resp.Builds {
			fmt.Println(*r.BuildId)
		}
	}

	return nil
}
