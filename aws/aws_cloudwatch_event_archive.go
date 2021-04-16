// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudwatchEventArchive(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Cloudwatcheventsconn.ListArchives(ctx, &cloudwatchevents.ListArchivesInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.Archives) > 0 {

		for _, r := range resp.Archives {

			t := *r.CreationTime
			result = append(result, terraform.Resource{
				Type:      "aws_cloudwatch_event_archive",
				ID:        *r.ArchiveName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
