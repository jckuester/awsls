// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func ListRoute53Zone(client *Client) error {
	req := client.route53conn.ListHostedZonesRequest(&route53.ListHostedZonesInput{})

	p := route53.NewListHostedZonesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.HostedZones {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
