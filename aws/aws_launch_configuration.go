// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLaunchConfiguration(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(client.Autoscalingconn, &autoscaling.DescribeLaunchConfigurationsInput{})
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range resp.LaunchConfigurations {

			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_launch_configuration",
				ID:        *r.LaunchConfigurationName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
