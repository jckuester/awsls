// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func ListSfnStateMachine(client *Client) error {
	req := client.sfnconn.ListStateMachinesRequest(&sfn.ListStateMachinesInput{})

	p := sfn.NewListStateMachinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.StateMachines {
			fmt.Println(*r.StateMachineArn)

			fmt.Printf("CreatedAt: %s\n", *r.CreationDate)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
