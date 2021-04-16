// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListApigatewayv2VpcLink(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Apigatewayv2conn.GetVpcLinks(ctx, &apigatewayv2.GetVpcLinksInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.Items) > 0 {

		for _, r := range resp.Items {

			result = append(result, terraform.Resource{
				Type:      "aws_apigatewayv2_vpc_link",
				ID:        *r.VpcLinkId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
