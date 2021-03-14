// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListImagebuilderImagePipeline(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Imagebuilderconn.ListImagePipelines(ctx, &imagebuilder.ListImagePipelinesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ImagePipelineList) > 0 {

		for _, r := range resp.ImagePipelineList {

			result = append(result, terraform.Resource{
				Type:      "aws_imagebuilder_image_pipeline",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
