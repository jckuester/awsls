// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigConfigurationRecorder(client *Client) {
	req := client.configserviceconn.DescribeConfigurationRecordersRequest(&configservice.DescribeConfigurationRecordersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_config_configuration_recorder: %s", err)
	} else {
		if len(resp.ConfigurationRecorders) > 0 {
			fmt.Println("")
			fmt.Println("aws_config_configuration_recorder:")
			for _, r := range resp.ConfigurationRecorders {
				fmt.Println(*r.Name)

			}
		}
	}

}
