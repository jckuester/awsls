// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListImagebuilderComponent(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Imagebuilderconn.ListComponentsRequest(&imagebuilder.ListComponentsInput{})

	var result []terraform.Resource

	p := imagebuilder.NewListComponentsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ComponentVersionList {

			result = append(result, terraform.Resource{
				Type:      "aws_imagebuilder_component",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
