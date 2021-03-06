// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListImagebuilderImage(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Imagebuilderconn.ListImagesRequest(&imagebuilder.ListImagesInput{})

	var result []terraform.Resource

	p := imagebuilder.NewListImagesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.ImageVersionList {

			result = append(result, terraform.Resource{
				Type:      "aws_imagebuilder_image",
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
