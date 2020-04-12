// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func ListRoute53HealthCheck(client *Client) {
	req := client.route53conn.ListHealthChecksRequest(&route53.ListHealthChecksInput{})

	p := route53.NewListHealthChecksPaginator(req)
	fmt.Println("")
	fmt.Println("aws_route53_health_check:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.HealthChecks {
			fmt.Println(*r.Id)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_route53_health_check: %s", err)
	}

}
