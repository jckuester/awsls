// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/worklink"
)

func ListWorklinkFleet(client *Client) ([]Resource, error) {
	req := client.Worklinkconn.ListFleetsRequest(&worklink.ListFleetsInput{})

	var result []Resource

	p := worklink.NewListFleetsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.FleetSummaryList {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}
			t := *r.CreatedTime
			result = append(result, Resource{
				Type:      "aws_worklink_fleet",
				ID:        *r.FleetArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
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
