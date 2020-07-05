// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ListSsmActivation(client *Client) ([]Resource, error) {
	req := client.Ssmconn.DescribeActivationsRequest(&ssm.DescribeActivationsInput{})

	var result []Resource

	p := ssm.NewDescribeActivationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ActivationList {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_ssm_activation",
				ID:     *r.ActivationId,
				Region: client.Ssmconn.Config.Region,
				Tags:   tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
