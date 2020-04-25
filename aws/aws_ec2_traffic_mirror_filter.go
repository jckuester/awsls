// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TrafficMirrorFilter(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeTrafficMirrorFiltersRequest(&ec2.DescribeTrafficMirrorFiltersInput{})

	var result []Resource

	p := ec2.NewDescribeTrafficMirrorFiltersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TrafficMirrorFilters {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_ec2_traffic_mirror_filter",
				ID:     *r.TrafficMirrorFilterId,
				Region: client.ec2conn.Config.Region,
				Tags:   tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
