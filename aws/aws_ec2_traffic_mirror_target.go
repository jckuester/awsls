// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2TrafficMirrorTarget(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeTrafficMirrorTargetsRequest(&ec2.DescribeTrafficMirrorTargetsInput{})

	var result []Resource

	p := ec2.NewDescribeTrafficMirrorTargetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.TrafficMirrorTargets {
			if *r.OwnerId != client.AccountID {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:      "aws_ec2_traffic_mirror_target",
				ID:        *r.TrafficMirrorTargetId,
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
