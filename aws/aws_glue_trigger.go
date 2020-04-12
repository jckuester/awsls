// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueTrigger(client *Client) error {
	req := client.glueconn.GetTriggersRequest(&glue.GetTriggersInput{})

	p := glue.NewGetTriggersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Triggers {
			fmt.Println(*r.Name)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
