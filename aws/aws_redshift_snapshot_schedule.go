// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListRedshiftSnapshotSchedule(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := redshift.NewDescribeSnapshotSchedulesPaginator(client.Redshiftconn, &redshift.DescribeSnapshotSchedulesInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.SnapshotSchedules {

			result = append(result, terraform.Resource{
				Type:      "aws_redshift_snapshot_schedule",
				ID:        *r.ScheduleIdentifier,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
