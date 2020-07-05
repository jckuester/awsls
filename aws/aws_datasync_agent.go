// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

func ListDatasyncAgent(client *Client) ([]Resource, error) {
	req := client.Datasyncconn.ListAgentsRequest(&datasync.ListAgentsInput{})

	var result []Resource

	p := datasync.NewListAgentsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Agents {

			result = append(result, Resource{
				Type:   "aws_datasync_agent",
				ID:     *r.AgentArn,
				Region: client.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
