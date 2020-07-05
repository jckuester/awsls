// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListPlacementGroup(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribePlacementGroupsRequest(&ec2.DescribePlacementGroupsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.PlacementGroups) > 0 {
		for _, r := range resp.PlacementGroups {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_placement_group",
				ID:     *r.GroupName,
				Region: client.Ec2conn.Config.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
