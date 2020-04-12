// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

func ListDatasyncTask(client *Client) error {
	req := client.datasyncconn.ListTasksRequest(&datasync.ListTasksInput{})

	p := datasync.NewListTasksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Tasks {
			fmt.Println(*r.TaskArn)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
