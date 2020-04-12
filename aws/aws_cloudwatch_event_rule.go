// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func ListCloudwatchEventRule(client *Client) {
	req := client.cloudwatcheventsconn.ListRulesRequest(&cloudwatchevents.ListRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_cloudwatch_event_rule: %s", err)
	} else {
		if len(resp.Rules) > 0 {
			fmt.Println("")
			fmt.Println("aws_cloudwatch_event_rule:")
			for _, r := range resp.Rules {
				fmt.Println(*r.Name)

			}
		}
	}

}
