// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

func ListServiceDiscoveryService(client *Client) error {
	req := client.servicediscoveryconn.ListServicesRequest(&servicediscovery.ListServicesInput{})

	p := servicediscovery.NewListServicesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Services {
			fmt.Println(*r.Id)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
