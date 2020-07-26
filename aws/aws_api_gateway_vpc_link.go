// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func ListApiGatewayVpcLink(client *Client) ([]Resource, error) {
	req := client.Apigatewayconn.GetVpcLinksRequest(&apigateway.GetVpcLinksInput{})

	var result []Resource

	p := apigateway.NewGetVpcLinksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Items {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type:    "aws_api_gateway_vpc_link",
				ID:      *r.Id,
				Profile: client.Profile,
				Region:  client.Region,
				Tags:    tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
