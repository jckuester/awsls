// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListLaunchTemplate(client *Client) ([]Resource, error) {
	req := client.Ec2conn.DescribeLaunchTemplatesRequest(&ec2.DescribeLaunchTemplatesInput{})

	var result []Resource

	p := ec2.NewDescribeLaunchTemplatesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LaunchTemplates {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreateTime
			result = append(result, Resource{
				Type:      "aws_launch_template",
				ID:        *r.LaunchTemplateId,
				Profile:   client.Profile,
				Region:    client.Region,
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
