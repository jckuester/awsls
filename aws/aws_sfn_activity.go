// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func ListSfnActivity(client *Client) error {
	req := client.sfnconn.ListActivitiesRequest(&sfn.ListActivitiesInput{})

	p := sfn.NewListActivitiesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Activities {
			fmt.Println(*r.ActivityArn)

			fmt.Printf("CreatedAt: %s\n", *r.CreationDate)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
