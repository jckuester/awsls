// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueTrigger(client *Client) {
	req := client.glueconn.GetTriggersRequest(&glue.GetTriggersInput{})

	p := glue.NewGetTriggersPaginator(req)
	fmt.Println("")
	fmt.Println("aws_glue_trigger:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Triggers {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_glue_trigger: %s", err)
	}

}
