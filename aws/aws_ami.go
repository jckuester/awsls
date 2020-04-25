// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListAmi(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeImagesRequest(&ec2.DescribeImagesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Images) > 0 {
		for _, r := range resp.Images {
			if *r.OwnerId != client.accountid {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.CreationDate)
			if err != nil {
				return nil, err
			}
			result = append(result, Resource{
				Type:      "aws_ami",
				ID:        *r.ImageId,
				Region:    client.ec2conn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
