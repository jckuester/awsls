// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func ListEfsAccessPoint(client *Client) ([]Resource, error) {
	req := client.Efsconn.DescribeAccessPointsRequest(&efs.DescribeAccessPointsInput{})

	var result []Resource

	p := efs.NewDescribeAccessPointsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.AccessPoints {
			if *r.OwnerId != client.AccountID {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:      "aws_efs_access_point",
				ID:        *r.AccessPointId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
