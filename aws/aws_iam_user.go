// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamUser(client *Client) ([]Resource, error) {
	req := client.iamconn.ListUsersRequest(&iam.ListUsersInput{})

	var result []Resource

	p := iam.NewListUsersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Users {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreateDate
			result = append(result, Resource{
				Type:      "aws_iam_user",
				ID:        *r.UserName,
				Region:    client.iamconn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
