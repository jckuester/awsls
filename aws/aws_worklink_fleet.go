// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/worklink"
)

func ListWorklinkFleet(client *Client) ([]Resource, error) {
	req := client.worklinkconn.ListFleetsRequest(&worklink.ListFleetsInput{})

	var result []Resource

	p := worklink.NewListFleetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.FleetSummaryList {

			t := *r.CreatedTime
			result = append(result, Resource{
				Type: "aws_worklink_fleet",
				ID:   *r.FleetArn,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
