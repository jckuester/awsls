// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
)

func ListXrayGroup(client *Client) ([]Resource, error) {
	req := client.Xrayconn.GetGroupsRequest(&xray.GetGroupsInput{})

	var result []Resource

	p := xray.NewGetGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Groups {

			result = append(result, Resource{
				Type:      "aws_xray_group",
				ID:        *r.GroupARN,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
