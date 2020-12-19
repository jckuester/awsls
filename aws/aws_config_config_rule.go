// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigConfigRule(client *Client) ([]Resource, error) {
	req := client.Configserviceconn.DescribeConfigRulesRequest(&configservice.DescribeConfigRulesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ConfigRules) > 0 {

		for _, r := range resp.ConfigRules {

			result = append(result, Resource{
				Type:      "aws_config_config_rule",
				ID:        *r.ConfigRuleName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
