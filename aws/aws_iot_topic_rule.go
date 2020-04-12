// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotTopicRule(client *Client) {
	req := client.iotconn.ListTopicRulesRequest(&iot.ListTopicRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_iot_topic_rule: %s", err)
	} else {
		if len(resp.Rules) > 0 {
			fmt.Println("")
			fmt.Println("aws_iot_topic_rule:")
			for _, r := range resp.Rules {
				fmt.Println(*r.RuleName)

			}
		}
	}

}
