// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func ListCloudformationStackSet(client *Client) ([]Resource, error) {
	req := client.cloudformationconn.ListStackSetsRequest(&cloudformation.ListStackSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Summaries) > 0 {
		for _, r := range resp.Summaries {

			result = append(result, Resource{
				Type: "aws_cloudformation_stack_set",
				ID:   *r.StackSetName,
			})
		}
	}

	return result, nil
}
