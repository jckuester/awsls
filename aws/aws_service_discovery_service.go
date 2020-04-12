// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

func ListServiceDiscoveryService(client *Client) {
	req := client.servicediscoveryconn.ListServicesRequest(&servicediscovery.ListServicesInput{})

	p := servicediscovery.NewListServicesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_service_discovery_service:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Services {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_service_discovery_service: %s", err)
	}

}
