// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftAlias(client *Client) {
	req := client.gameliftconn.ListAliasesRequest(&gamelift.ListAliasesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_gamelift_alias: %s", err)
	} else {
		if len(resp.Aliases) > 0 {
			fmt.Println("")
			fmt.Println("aws_gamelift_alias:")
			for _, r := range resp.Aliases {
				fmt.Println(*r.AliasId)

			}
		}
	}

}
