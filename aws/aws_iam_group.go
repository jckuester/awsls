// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListIamGroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Iamconn.ListGroupsRequest(&iam.ListGroupsInput{})

	var result []terraform.Resource

	p := iam.NewListGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

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

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
