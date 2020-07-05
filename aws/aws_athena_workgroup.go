// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/athena"
)

func ListAthenaWorkgroup(client *Client) ([]Resource, error) {
	req := client.Athenaconn.ListWorkGroupsRequest(&athena.ListWorkGroupsInput{})

	var result []Resource

	p := athena.NewListWorkGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.WorkGroups {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:   "aws_athena_workgroup",
				ID:     *r.Name,
				Region: client.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
