// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func ListRoute53HealthCheck(client *Client) error {
	req := client.route53conn.ListHealthChecksRequest(&route53.ListHealthChecksInput{})

	p := route53.NewListHealthChecksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.HealthChecks {
			fmt.Println(*r.Id)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
