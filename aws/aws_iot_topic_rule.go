// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func ListIotTopicRule(client *Client) ([]Resource, error) {
	req := client.Iotconn.ListTopicRulesRequest(&iot.ListTopicRulesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) > 0 {
		for _, r := range resp.Rules {

			result = append(result, Resource{
				Type:   "aws_iot_topic_rule",
				ID:     *r.RuleName,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
