// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftBuild(client *Client) ([]Resource, error) {
	req := client.gameliftconn.ListBuildsRequest(&gamelift.ListBuildsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Builds) > 0 {
		for _, r := range resp.Builds {

			t := *r.CreationTime
			result = append(result, Resource{
				Type: "aws_gamelift_build",
				ID:   *r.BuildId,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
