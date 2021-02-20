// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDatasyncAgent(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Datasyncconn.ListAgentsRequest(&datasync.ListAgentsInput{})

	var result []terraform.Resource

	p := datasync.NewListAgentsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Agents {

			result = append(result, terraform.Resource{
				Type:      "aws_datasync_agent",
				ID:        *r.AgentArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
