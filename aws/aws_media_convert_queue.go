// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

func ListMediaConvertQueue(client *Client) {
	req := client.mediaconvertconn.ListQueuesRequest(&mediaconvert.ListQueuesInput{})

	p := mediaconvert.NewListQueuesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_media_convert_queue:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Queues {
			fmt.Println(*r.Name)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_media_convert_queue: %s", err)
	}

}
