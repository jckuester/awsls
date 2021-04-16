// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListImagebuilderImageRecipe(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Imagebuilderconn.ListImageRecipes(ctx, &imagebuilder.ListImageRecipesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.ImageRecipeSummaryList) > 0 {

		for _, r := range resp.ImageRecipeSummaryList {

			result = append(result, terraform.Resource{
				Type:      "aws_imagebuilder_image_recipe",
				ID:        *r.Arn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
