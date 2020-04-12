// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

func ListDatasyncAgent(client *Client) error {
	req := client.datasyncconn.ListAgentsRequest(&datasync.ListAgentsInput{})

	p := datasync.NewListAgentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Agents {
			fmt.Println(*r.AgentArn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
