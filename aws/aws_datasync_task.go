// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

func ListDatasyncTask(client *Client) {
	req := client.datasyncconn.ListTasksRequest(&datasync.ListTasksInput{})

	p := datasync.NewListTasksPaginator(req)
	fmt.Println("")
	fmt.Println("aws_datasync_task:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Tasks {
			fmt.Println(*r.TaskArn)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_datasync_task: %s", err)
	}

}
