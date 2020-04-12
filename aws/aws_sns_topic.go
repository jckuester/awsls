// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsTopic(client *Client) {
	req := client.snsconn.ListTopicsRequest(&sns.ListTopicsInput{})

	p := sns.NewListTopicsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_sns_topic:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Topics {
			fmt.Println(*r.TopicArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_sns_topic: %s", err)
	}

}
