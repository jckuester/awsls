// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func ListApiGatewayUsagePlan(client *Client) ([]Resource, error) {
	req := client.apigatewayconn.GetUsagePlansRequest(&apigateway.GetUsagePlansInput{})

	var result []Resource

	p := apigateway.NewGetUsagePlansPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Items {
			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type: "aws_api_gateway_usage_plan",
				ID:   *r.Id,
				Tags: tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
