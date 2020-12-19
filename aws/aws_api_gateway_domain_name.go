// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func ListApiGatewayDomainName(client *Client) ([]Resource, error) {
	req := client.Apigatewayconn.GetDomainNamesRequest(&apigateway.GetDomainNamesInput{})

	var result []Resource

	p := apigateway.NewGetDomainNamesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Items {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type:      "aws_api_gateway_domain_name",
				ID:        *r.DomainName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
