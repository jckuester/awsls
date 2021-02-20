// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListConfigConfigurationRecorder(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Configserviceconn.DescribeConfigurationRecordersRequest(&configservice.DescribeConfigurationRecordersInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ConfigurationRecorders) > 0 {

		for _, r := range resp.ConfigurationRecorders {

			result = append(result, terraform.Resource{
				Type:      "aws_config_configuration_recorder",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
