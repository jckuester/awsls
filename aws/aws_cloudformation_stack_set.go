// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func ListCloudformationStackSet(client *Client) {
	req := client.cloudformationconn.ListStackSetsRequest(&cloudformation.ListStackSetsInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_cloudformation_stack_set: %s", err)
	} else {
		if len(resp.Summaries) > 0 {
			fmt.Println("")
			fmt.Println("aws_cloudformation_stack_set:")
			for _, r := range resp.Summaries {
				fmt.Println(*r.StackSetName)

			}
		}
	}

}
