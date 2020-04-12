// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

func ListDmsEndpoint(client *Client) {
	req := client.databasemigrationserviceconn.DescribeEndpointsRequest(&databasemigrationservice.DescribeEndpointsInput{})

	p := databasemigrationservice.NewDescribeEndpointsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_dms_endpoint:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Endpoints {
			fmt.Println(*r.EndpointIdentifier)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_dms_endpoint: %s", err)
	}

}
