// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigConfigRule(client *Client) error {
	req := client.configserviceconn.DescribeConfigRulesRequest(&configservice.DescribeConfigRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		return err
	}

	if len(resp.ConfigRules) > 0 {
		for _, r := range resp.ConfigRules {
			fmt.Println(*r.ConfigRuleName)

		}
	}

	return nil
}
