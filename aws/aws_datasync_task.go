// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/datasync"
)

func ListDatasyncTask(client *Client) ([]Resource, error) {
	req := client.Datasyncconn.ListTasksRequest(&datasync.ListTasksInput{})

	var result []Resource

	p := datasync.NewListTasksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Tasks {

			result = append(result, Resource{
				Type:   "aws_datasync_task",
				ID:     *r.TaskArn,
				Region: client.Datasyncconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
