// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func ListSfnActivity(client *Client) ([]Resource, error) {
	req := client.Sfnconn.ListActivitiesRequest(&sfn.ListActivitiesInput{})

	var result []Resource

	p := sfn.NewListActivitiesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Activities {

			t := *r.CreationDate
			result = append(result, Resource{
				Type:    "aws_sfn_activity",
				ID:      *r.ActivityArn,
				Profile: client.Profile,
				Region:  client.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
