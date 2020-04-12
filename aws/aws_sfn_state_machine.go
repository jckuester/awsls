// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func ListSfnStateMachine(client *Client) {
	req := client.sfnconn.ListStateMachinesRequest(&sfn.ListStateMachinesInput{})

	p := sfn.NewListStateMachinesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_sfn_state_machine:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.StateMachines {
			fmt.Println(*r.StateMachineArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_sfn_state_machine: %s", err)
	}

}
