// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
)

func ListGameliftAlias(client *Client) ([]Resource, error) {
	req := client.Gameliftconn.ListAliasesRequest(&gamelift.ListAliasesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Aliases) > 0 {

		for _, r := range resp.Aliases {

			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_gamelift_alias",
				ID:        *r.AliasId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
