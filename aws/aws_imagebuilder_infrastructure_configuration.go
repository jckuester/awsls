// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
)

func ListImagebuilderInfrastructureConfiguration(client *Client) ([]Resource, error) {
	req := client.Imagebuilderconn.ListInfrastructureConfigurationsRequest(&imagebuilder.ListInfrastructureConfigurationsInput{})

	var result []Resource

	p := imagebuilder.NewListInfrastructureConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.InfrastructureConfigurationSummaryList {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}

			result = append(result, Resource{
				Type:      "aws_imagebuilder_infrastructure_configuration",
				ID:        *r.Arn,
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
