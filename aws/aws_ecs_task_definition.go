// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func ListEcsTaskDefinition(client *Client) ([]Resource, error) {
	req := client.Ecsconn.ListTaskDefinitionsRequest(&ecs.ListTaskDefinitionsInput{})

	var result []Resource

	p := ecs.NewListTaskDefinitionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.TaskDefinitionArns {

			result = append(result, Resource{
				Type:      "aws_ecs_task_definition",
				ID:        r,
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
