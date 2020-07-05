// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

func ListSagemakerEndpoint(client *Client) ([]Resource, error) {
	req := client.Sagemakerconn.ListEndpointsRequest(&sagemaker.ListEndpointsInput{})

	var result []Resource

	p := sagemaker.NewListEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Endpoints {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:   "aws_sagemaker_endpoint",
				ID:     *r.EndpointName,
				Region: client.Sagemakerconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
