// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSfnActivity(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Sfnconn.ListActivitiesRequest(&sfn.ListActivitiesInput{})

	var result []terraform.Resource

	p := sfn.NewListActivitiesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Activities {

			t := *r.CreationDate
			result = append(result, terraform.Resource{
				Type:      "aws_sfn_activity",
				ID:        *r.ActivityArn,
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
