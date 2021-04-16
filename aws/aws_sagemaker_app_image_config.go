// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSagemakerAppImageConfig(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := sagemaker.NewListAppImageConfigsPaginator(client.Sagemakerconn, &sagemaker.ListAppImageConfigsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.AppImageConfigs {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_sagemaker_app_image_config",
				ID:        *r.AppImageConfigName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
