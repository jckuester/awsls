// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigConfigRule(client *Client) {
	req := client.configserviceconn.DescribeConfigRulesRequest(&configservice.DescribeConfigRulesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_config_config_rule: %s", err)
	} else {
		if len(resp.ConfigRules) > 0 {
			fmt.Println("")
			fmt.Println("aws_config_config_rule:")
			for _, r := range resp.ConfigRules {
				fmt.Println(*r.ConfigRuleName)

			}
		}
	}

}
