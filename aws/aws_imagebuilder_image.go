// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListImagebuilderImage(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Imagebuilderconn.ListImages(ctx, &imagebuilder.ListImagesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ImageVersionList) > 0 {

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

	return result, nil
}
