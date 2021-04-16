// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListIamGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := iam.NewListGroupsPaginator(client.Iamconn, &iam.ListGroupsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Groups {

			t := *r.CreateDate
			result = append(result, terraform.Resource{
				Type:      "aws_iam_group",
				ID:        *r.GroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
