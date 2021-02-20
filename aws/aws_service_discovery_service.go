// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListServiceDiscoveryService(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Servicediscoveryconn.ListServicesRequest(&servicediscovery.ListServicesInput{})

	var result []terraform.Resource

	p := servicediscovery.NewListServicesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Services {

			t := *r.CreateDate
			result = append(result, terraform.Resource{
				Type:      "aws_service_discovery_service",
				ID:        *r.Id,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
