// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsEndpoint(client *Client) ([]Resource, error) {
	req := client.Databasemigrationserviceconn.DescribeEndpointsRequest(&databasemigrationservice.DescribeEndpointsInput{})

	var result []Resource

	p := databasemigrationservice.NewDescribeEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Endpoints {

			result = append(result, Resource{
				Type:      "aws_dms_endpoint",
				ID:        *r.EndpointIdentifier,
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
