// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"
)

func ListWafv2WebAclLoggingConfiguration(client *Client) ([]Resource, error) {
	req := client.wafv2conn.ListLoggingConfigurationsRequest(&wafv2.ListLoggingConfigurationsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.LoggingConfigurations) > 0 {
		for _, r := range resp.LoggingConfigurations {

			result = append(result, Resource{
				Type:   "aws_wafv2_web_acl_logging_configuration",
				ID:     *r.ResourceArn,
				Region: client.wafv2conn.Config.Region,
			})
		}
	}

	return result, nil
}
