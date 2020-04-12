// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func ListLaunchConfiguration(client *Client) {
	req := client.autoscalingconn.DescribeLaunchConfigurationsRequest(&autoscaling.DescribeLaunchConfigurationsInput{})

	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(req)
	fmt.Println("")
	fmt.Println("aws_launch_configuration:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.LaunchConfigurations {
			fmt.Println(*r.LaunchConfigurationName)

		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_launch_configuration: %s", err)
	}

}
