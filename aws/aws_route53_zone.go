// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func ListRoute53Zone(client *Client) {
	req := client.route53conn.ListHostedZonesRequest(&route53.ListHostedZonesInput{})

	p := route53.NewListHostedZonesPaginator(req)
	fmt.Println("")
	fmt.Println("aws_route53_zone:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.HostedZones {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_route53_zone: %s", err)
	}

}
