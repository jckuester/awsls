// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func ListGlueSecurityConfiguration(client *Client) ([]Resource, error) {
	req := client.Glueconn.GetSecurityConfigurationsRequest(&glue.GetSecurityConfigurationsInput{})

	var result []Resource

	p := glue.NewGetSecurityConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.SecurityConfigurations {

			result = append(result, Resource{
				Type:   "aws_glue_security_configuration",
				ID:     *r.Name,
				Region: client.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
