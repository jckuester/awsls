// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

func ListServiceDiscoveryService(client *Client) ([]Resource, error) {
	req := client.servicediscoveryconn.ListServicesRequest(&servicediscovery.ListServicesInput{})

	var result []Resource

	p := servicediscovery.NewListServicesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Services {

			t := *r.CreateDate
			result = append(result, Resource{
				Type: "aws_service_discovery_service",
				ID:   *r.Id,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
