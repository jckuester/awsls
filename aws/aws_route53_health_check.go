// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func ListRoute53HealthCheck(client *Client) ([]Resource, error) {
	req := client.Route53conn.ListHealthChecksRequest(&route53.ListHealthChecksInput{})

	var result []Resource

	p := route53.NewListHealthChecksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.HealthChecks {

			result = append(result, Resource{
				Type:   "aws_route53_health_check",
				ID:     *r.Id,
				Region: client.Route53conn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
