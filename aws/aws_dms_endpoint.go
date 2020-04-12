// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsEndpoint(client *Client) error {
	req := client.databasemigrationserviceconn.DescribeEndpointsRequest(&databasemigrationservice.DescribeEndpointsInput{})

	p := databasemigrationservice.NewDescribeEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Endpoints {
			fmt.Println(*r.EndpointIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
