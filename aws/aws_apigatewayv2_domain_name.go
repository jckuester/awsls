// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListApigatewayv2DomainName(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Apigatewayv2conn.GetDomainNames(ctx, &apigatewayv2.GetDomainNamesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.Items) > 0 {

		for _, r := range resp.Items {

			result = append(result, terraform.Resource{
				Type:      "aws_apigatewayv2_domain_name",
				ID:        *r.DomainName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
