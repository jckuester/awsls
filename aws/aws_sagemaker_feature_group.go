// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSagemakerFeatureGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Sagemakerconn.ListFeatureGroups(ctx, &sagemaker.ListFeatureGroupsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.FeatureGroupSummaries) > 0 {

		for _, r := range resp.FeatureGroupSummaries {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_sagemaker_feature_group",
				ID:        *r.FeatureGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
