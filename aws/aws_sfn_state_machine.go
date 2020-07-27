// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func ListSfnStateMachine(client *Client) ([]Resource, error) {
	req := client.Sfnconn.ListStateMachinesRequest(&sfn.ListStateMachinesInput{})

	var result []Resource

	p := sfn.NewListStateMachinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.StateMachines {

			t := *r.CreationDate
			result = append(result, Resource{
				Type:      "aws_sfn_state_machine",
				ID:        *r.StateMachineArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
