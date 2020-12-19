// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListSnsTopic(client *Client) ([]Resource, error) {
	req := client.Snsconn.ListTopicsRequest(&sns.ListTopicsInput{})

	var result []Resource

	p := sns.NewListTopicsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Topics {

			result = append(result, Resource{
				Type:      "aws_sns_topic",
				ID:        *r.TopicArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
