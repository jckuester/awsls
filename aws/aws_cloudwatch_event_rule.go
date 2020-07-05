// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func ListCloudwatchEventRule(client *Client) ([]Resource, error) {
	req := client.Cloudwatcheventsconn.ListRulesRequest(&cloudwatchevents.ListRulesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) > 0 {
		for _, r := range resp.Rules {

			result = append(result, Resource{
				Type:   "aws_cloudwatch_event_rule",
				ID:     *r.Name,
				Region: client.Region,
			})
		}
	}

	return result, nil
}
