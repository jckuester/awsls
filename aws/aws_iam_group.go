// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamGroup(client *Client) ([]Resource, error) {
	req := client.Iamconn.ListGroupsRequest(&iam.ListGroupsInput{})

	var result []Resource

	p := iam.NewListGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Groups {

			t := *r.CreateDate
			result = append(result, Resource{
				Type:      "aws_iam_group",
				ID:        *r.GroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
