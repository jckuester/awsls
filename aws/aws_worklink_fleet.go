// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWorklinkFleet(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := worklink.NewListFleetsPaginator(client.Worklinkconn, &worklink.ListFleetsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.FleetSummaryList {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}
			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_worklink_fleet",
				ID:        *r.FleetArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
