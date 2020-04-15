// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emr"
)

func ListEmrSecurityConfiguration(client *Client) ([]Resource, error) {
	req := client.emrconn.ListSecurityConfigurationsRequest(&emr.ListSecurityConfigurationsInput{})

	var result []Resource

	p := emr.NewListSecurityConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecurityConfigurations {

			result = append(result, Resource{
				Type: "aws_emr_security_configuration",
				ID:   *r.Name,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
