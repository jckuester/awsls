// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

func ListDatasyncAgent(client *Client) {
	req := client.datasyncconn.ListAgentsRequest(&datasync.ListAgentsInput{})

	p := datasync.NewListAgentsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_datasync_agent:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Agents {
			fmt.Println(*r.AgentArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_datasync_agent: %s", err)
	}

}
