// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func ListLaunchConfiguration(client *Client) error {
	req := client.autoscalingconn.DescribeLaunchConfigurationsRequest(&autoscaling.DescribeLaunchConfigurationsInput{})

	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LaunchConfigurations {
			fmt.Println(*r.LaunchConfigurationName)

			fmt.Printf("CreatedAt: %s\n", *r.CreatedTime)
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
