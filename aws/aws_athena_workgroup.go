// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListAthenaWorkgroup(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Athenaconn.ListWorkGroupsRequest(&athena.ListWorkGroupsInput{})

	var result []terraform.Resource

	p := athena.NewListWorkGroupsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.WorkGroups {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_athena_workgroup",
				ID:        *r.Name,
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
