// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftAlias(client *Client) error {
	req := client.gameliftconn.ListAliasesRequest(&gamelift.ListAliasesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.Aliases) > 0 {
		for _, r := range resp.Aliases {
			fmt.Println(*r.AliasId)

			fmt.Printf("CreatedAt: %s\n", *r.CreationTime)
		}
	}

	return nil
}
