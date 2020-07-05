// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

func ListApigatewayv2VpcLink(client *Client) ([]Resource, error) {
	req := client.Apigatewayv2conn.GetVpcLinksRequest(&apigatewayv2.GetVpcLinksInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Items) > 0 {
		for _, r := range resp.Items {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type:   "aws_apigatewayv2_vpc_link",
				ID:     *r.VpcLinkId,
				Region: client.Region,
				Tags:   tags,
			})
		}
	}

	return result, nil
}
