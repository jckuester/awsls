// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsTopic(client *Client) error {
	req := client.snsconn.ListTopicsRequest(&sns.ListTopicsInput{})

	p := sns.NewListTopicsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Topics {
			fmt.Println(*r.TopicArn)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
