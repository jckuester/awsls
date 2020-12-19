// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

func ListCloudtrail(client *Client) ([]Resource, error) {
	req := client.Cloudtrailconn.DescribeTrailsRequest(&cloudtrail.DescribeTrailsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.TrailList) > 0 {

		for _, r := range resp.TrailList {

			result = append(result, Resource{
				Type:      "aws_cloudtrail",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
