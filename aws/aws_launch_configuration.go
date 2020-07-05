// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func ListLaunchConfiguration(client *Client) ([]Resource, error) {
	req := client.Autoscalingconn.DescribeLaunchConfigurationsRequest(&autoscaling.DescribeLaunchConfigurationsInput{})

	var result []Resource

	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LaunchConfigurations {

			t := *r.CreatedTime
			result = append(result, Resource{
				Type:   "aws_launch_configuration",
				ID:     *r.LaunchConfigurationName,
				Region: client.Autoscalingconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
