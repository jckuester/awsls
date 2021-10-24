// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSagemakerModelPackageGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := sagemaker.NewListModelPackageGroupsPaginator(client.Sagemakerconn, &sagemaker.ListModelPackageGroupsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.ModelPackageGroupSummaryList {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_sagemaker_model_package_group",
				ID:        *r.ModelPackageGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
