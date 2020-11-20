// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueDevEndpoint(client *Client) ([]Resource, error) {
	req := client.Glueconn.GetDevEndpointsRequest(&glue.GetDevEndpointsInput{})

	var result []Resource

	p := glue.NewGetDevEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.DevEndpoints {

			result = append(result, Resource{
				Type:      "aws_glue_dev_endpoint",
				ID:        *r.EndpointName,
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
