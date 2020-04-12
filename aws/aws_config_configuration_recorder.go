// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigConfigurationRecorder(client *Client) error {
	req := client.configserviceconn.DescribeConfigurationRecordersRequest(&configservice.DescribeConfigurationRecordersInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.ConfigurationRecorders) > 0 {
		for _, r := range resp.ConfigurationRecorders {
			fmt.Println(*r.Name)
		}
	}

	return nil
}
