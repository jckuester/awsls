// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListPlacementGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Ec2conn.DescribePlacementGroups(ctx, &ec2.DescribePlacementGroupsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.PlacementGroups) > 0 {

		for _, r := range resp.PlacementGroups {

			result = append(result, terraform.Resource{
				Type:      "aws_placement_group",
				ID:        *r.GroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
