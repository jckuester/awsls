// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

func ListMediaConvertQueue(client *Client) error {
	req := client.mediaconvertconn.ListQueuesRequest(&mediaconvert.ListQueuesInput{})

	p := mediaconvert.NewListQueuesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Queues {
			fmt.Println(*r.Name)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
