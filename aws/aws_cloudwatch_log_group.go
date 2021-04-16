// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCloudwatchLogGroup(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := cloudwatchlogs.NewDescribeLogGroupsPaginator(client.Cloudwatchlogsconn, &cloudwatchlogs.DescribeLogGroupsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.LogGroups {

			t := time.Unix(0, *r.CreationTime*1000000).UTC()
			result = append(result, terraform.Resource{
				Type:      "aws_cloudwatch_log_group",
				ID:        *r.LogGroupName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
